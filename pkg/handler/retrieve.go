package handler

import (
	"errors"
	"fmt"
	"sso/database"
	"sso/model"
	"sso/utils"
)

type Retrieve interface {
	SendEmailCode(email string) error
	Retrieve(password string) error
}

func NewRetrieveHandler(operateId string) Retrieve {
	return &retrieveHandler{OperateId: operateId}
}

type retrieveHandler struct {
	registerHandler
	OperateId string
}

// SendEmailCode 发送邮箱验证码
func (r *retrieveHandler) SendEmailCode(email string) error {
	// 1. 校验邮箱地址是否存在
	if e := (&model.TUser{}).FirstByEmail(email); e != nil {
		return e
	}
	// 2. 校验验证码是否存在
	k := fmt.Sprintf("%s-%s", r.OperateId, email)
	if database.R.HExists(k, codeField).Val() {
		return fmt.Errorf("验证码还在有效期，请稍候重试")
	}
	// 3. 生成随机验证码
	code := r.geneCode()
	if e := utils.SendEmail("SSO平台验证邮箱", r.template(code), "", email); e != nil {
		return fmt.Errorf("发送失败: %w", e)
	}
	// 4. 保存邮箱验证码
	if e := r.hSetEmailCode(email, code); e != nil {
		return fmt.Errorf("保存验证码失败: %w", e)
	}
	return nil
}

// Retrieve 找回密码
func (r *retrieveHandler) Retrieve(password string) error {
	// 1. 校验邮箱是否验证通过
	if !r.verifyEmailPass() {
		return errors.New("请先验证邮箱信息")
	}
	// 2. 获取私钥解密密码
	pwd, e := r.decodePwd(password)
	if e != nil {
		return e
	}
	// 3. 更新密码信息
	email, e := r.hGetEmail()
	if e != nil {
		return fmt.Errorf("邮箱获取失败: %w", e)
	}
	user := model.TUser{}

	if e := user.FirstByEmail(email); e != nil {
		return e
	}
	// hash保存密码
	if e := database.DB.Model(&user).Update("pwd", utils.HashString(pwd)).Error; e != nil {
		return fmt.Errorf("密码重置失败: %w", e)
	}
	return nil
}
