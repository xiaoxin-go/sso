package handler

import (
	"bytes"
	"fmt"
	"github.com/pquerna/otp/totp"
	"go.uber.org/zap"
	"image/png"
	"math/rand"
	"regexp"
	"sso/database"
	"sso/model"
	"sso/utils"
	"strconv"
	"time"
)

const (
	codeField    = "code"
	emailField   = "email"
	secretField  = "secret"
	otpPassField = "otpPass"
)

type Register interface {
	SendEmailCode(email string) error
	VerifyEmailCode(email, code string) error
	GetOtpQrCode() ([]byte, error)
	VerifyOtpCode(code uint32) error
	Register(user *model.TUser) error
}

func NewRegisterHandler(operateId string) Register {
	return &registerHandler{OperateId: operateId}
}

type registerHandler struct {
	OperateId string
	Err       error
}

// SendEmailCode 发送验证码
func (r *registerHandler) SendEmailCode(email string) error {
	// 校验邮箱地址是否被占用
	if e := r.isExistsEmail(email); e != nil {
		return e
	}
	k := fmt.Sprintf("%s-%s", r.OperateId, email)
	// 从redis中获取验证码, 如果验证码有效，则不需要发送
	if database.R.HExists(k, codeField).Val() {
		return fmt.Errorf("验证码还在有效期，请稍候重试")
	}

	// 生成验证码
	code := r.geneCode()
	// 发送验证码
	if e := utils.SendEmail("SSO平台验证邮箱", r.template(code), "", email); e != nil {
		return fmt.Errorf("发送失败: %w", e)
	}

	// 保存验证码
	if e := r.hSetEmailCode(email, code); e != nil {
		return e
	}
	return nil
}

// VerifyEmailCode 校验验证码
func (r *registerHandler) VerifyEmailCode(email, code string) error {
	// 校验是否正确
	if !r.verifyCode(email, code) {
		return fmt.Errorf("验证码不正确")
	}
	// 保存校验结果
	if e := r.hSetEmail(email); e != nil {
		return e
	}
	return nil
}

// GetOtpQrCode 获取otp二维码
func (r *registerHandler) GetOtpQrCode() ([]byte, error) {
	// 获取邮箱是否验证通过
	email, e := r.hGetEmail()
	if e != nil {
		return nil, fmt.Errorf("邮箱验证过期")
	}
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      email,
		AccountName: email,
	})
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		return nil, fmt.Errorf("生成二维码失败: %w", err)
	}
	if err := png.Encode(&buf, img); err != nil {
		return nil, fmt.Errorf("生成二维码失败: %w", err)
	}
	// 保存密钥
	if e := r.hSetSecret(key.Secret()); e != nil {
		return nil, fmt.Errorf("保存密钥失败: %w", e)
	}
	return buf.Bytes(), nil
}

// VerifyOtpCode 校验otp code
func (r *registerHandler) VerifyOtpCode(code uint32) error {
	// 获取Secret
	secret, err := r.hGetSecret()
	if err != nil {
		return fmt.Errorf("二维码失效，请获取重试")
	}
	// 校验otp
	str := strconv.Itoa(int(code))
	if !totp.Validate(str, secret) {
		return fmt.Errorf("otp不正确")
	}
	// 保存校验结果
	if e := r.hSetVerifyOtpResult(); e != nil {
		return fmt.Errorf("保存校验结果异常: %w, 请重试", e)
	}
	return nil
}

// Register 注册
func (r *registerHandler) Register(user *model.TUser) error {
	l := zap.L().With(zap.String("func", "Register"), zap.String("OperateId", r.OperateId),
		zap.String("username", user.Username))
	l.Debug("用户注册--->")
	if e := r.isExistsUsername(user.Username); e != nil {
		return e
	}
	// 校验otp是否验证通过
	if !r.verifyOtpPass() {
		return fmt.Errorf("otp未验证通过")
	}
	// 解密密码，校验密码安全性
	pwd, e := r.decodePwd(user.Password)
	if e != nil {
		return e
	}
	if !r.verifyPwd(pwd) {
		return fmt.Errorf("密码格式不正确")
	}
	// 从缓存中取出邮箱和otp secret
	email, e := r.hGetEmail()
	if e != nil {
		return fmt.Errorf("获取email失败: %w", e)
	}
	secret, e := r.hGetSecret()
	if e != nil {
		return fmt.Errorf("获取otp密钥失败: %w", e)
	}
	user.Email = email
	user.OtpSecret = secret
	// hash密码
	user.Password = utils.HashString(pwd)
	user.Enabled = 1
	// 保存用户
	if e := user.Create(); e != nil {
		return e
	}
	l.Debug("用户注册成功")
	return nil
}

// 校对密码格式
func (r *registerHandler) verifyPwd(pwd string) bool {
	// 密码不能小于8位
	if len(pwd) < 8 {
		return false
	}
	if ok, _ := regexp.MatchString("[a-z]", pwd); !ok {
		return false
	}
	if ok, _ := regexp.MatchString("[A-Z]", pwd); !ok {
		return false
	}
	if ok, _ := regexp.MatchString("[!@#$%^&*]", pwd); !ok {
		return false
	}
	return true
}

func (r *registerHandler) decodePwd(pwd string) (string, error) {
	privateKey, e := database.R.HGet(r.OperateId, "private_key").Result()
	if e != nil {
		return "", fmt.Errorf("密码私钥已过期，请重试")
	}
	result, e := utils.RsaDecrypt(pwd, privateKey)
	if e != nil {
		return "", fmt.Errorf("解密失败: %w", e)
	}
	return result, nil
}

func (r *registerHandler) verifyCode(email, code string) bool {
	// 从redis取出验证码并对比
	k := fmt.Sprintf("%s-%s", r.OperateId, email)
	return code == database.R.HGet(k, codeField).Val()
}

func (r *registerHandler) template(captcha string) string {
	tpl := `
	<h3>邮箱验证码: %s</h3>
	</br>
	<h3>请在5分钟内输入验证码</h3>`
	return fmt.Sprintf(tpl, captcha)
}

func (r *registerHandler) verifyEmailPass() bool {
	return database.R.HExists(r.OperateId, emailField).Val()
}

func (r *registerHandler) geneCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
func (r *registerHandler) hSetSecret(secret string) error {
	if e := database.R.HSet(r.OperateId, secretField, secret).Err(); e != nil {
		return fmt.Errorf("保存密钥失败: %w", e)
	}
	r.setOperateIdExpire()
	return nil
}
func (r *registerHandler) hGetSecret() (string, error) {
	return database.R.HGet(r.OperateId, secretField).Result()
}
func (r *registerHandler) hSetVerifyOtpResult() error {
	if e := database.R.HSet(r.OperateId, otpPassField, true).Err(); e != nil {
		return e
	}
	r.setOperateIdExpire()
	return nil
}
func (r *registerHandler) verifyOtpPass() bool {
	return database.R.HExists(r.OperateId, otpPassField).Val()
}
func (r *registerHandler) hSetEmail(email string) error {
	if e := database.R.HSet(r.OperateId, emailField, email).Err(); e != nil {
		return fmt.Errorf("保存邮箱失败: %w，请重试", e)
	}
	r.setOperateIdExpire()
	return nil
}
func (r *registerHandler) hSetEmailCode(email, code string) error {
	k := fmt.Sprintf("%s-%s", r.OperateId, email)
	// 保存验证码
	if e := database.R.HSet(k, codeField, code).Err(); e != nil {
		return fmt.Errorf("保存验证码异常: %w", e)
	}
	// 设置有效期
	r.setOperateIdExpire()
	return nil
}
func (r *registerHandler) setOperateIdExpire() {
	database.R.Expire(r.OperateId, 5*time.Minute)
}
func (r *registerHandler) hGetEmail() (string, error) {
	return database.R.HGet(r.OperateId, emailField).Result()
}
func (r *registerHandler) isExistsEmail(email string) error {
	if e := (&model.TUser{}).FirstByEmail(email); e == nil {
		return fmt.Errorf("邮箱地址已被占用")
	}
	return nil
}
func (r *registerHandler) isExistsUsername(username string) error {
	if e := (&model.TUser{}).FirstByUsername(username); e == nil {
		return fmt.Errorf("用户名已被占用")
	}
	return nil
}
