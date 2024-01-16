package handler

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"sso/database"
	"sso/model"
	"sso/utils"
	"time"
)

func NewLoginHandler(operateId string) *loginHandler {
	return &loginHandler{OperateId: operateId}
}

type loginHandler struct {
	OperateId string
}

func (h *loginHandler) Login(username, password string, otp uint32) (string, error) {
	l := zap.L().With(zap.String("func", "Login"), zap.String("OperateId", h.OperateId))
	l.Debug("用户登录")
	l.Debug("解密密码")
	pwd, e := h.decodePassword(password)
	if e != nil {
		l.Error("解密密码错误", zap.Error(e))
		return "", e
	}
	l.Debug("获取用户信息")
	user := model.TUser{}
	if e := user.FirstByNameOrEmail(username); e != nil {
		return "", e
	}

	l.Debug("验证密码")
	if user.Password != utils.HashString(pwd) {
		return "", fmt.Errorf("密码错误")
	}
	l.Debug("验证otp")
	if !utils.NewTotp(user.OtpSecret).Verify(otp) {
		return "", fmt.Errorf("otp过期")
	}
	l.Debug("生成session_id")
	sessionId := h.geneSessionId()
	if e := h.hSetSessionId(sessionId, &user); e != nil {
		return "", fmt.Errorf("保存session失败: %s", e.Error())
	}
	return sessionId, nil
}
func (h *loginHandler) hSetSessionId(sessionId string, user *model.TUser) error {
	if e := database.R.HSet(sessionId, "username", user.Username).Err(); e != nil {
		return e
	}
	if e := database.R.HSet(sessionId, "user_id", user.Id).Err(); e != nil {
		return e
	}
	database.R.Expire(sessionId, 24*time.Hour)
	return nil
}

func (h *loginHandler) decodePassword(password string) (string, error) {
	privateKey, e := database.R.HGet(h.OperateId, "private_key").Result()
	if e != nil {
		return "", fmt.Errorf("获取私钥异常: %w", e)
	}
	if privateKey == "" {
		return "", errors.New("密钥过期，请刷新页面重试")
	}
	result, e := utils.RsaDecrypt(password, privateKey)
	if e != nil {
		return "", fmt.Errorf("密码解密失败: %w", e)
	}
	return result, nil
}

func (h *loginHandler) geneSessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
