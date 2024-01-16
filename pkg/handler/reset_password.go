package handler

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"sso/database"
	"sso/model"
	"sso/utils"
	"time"
)

func NewResetPasswordHandler(operateId string) *resetPasswordHandler {
	h := &resetPasswordHandler{}
	h.OperateId = operateId
	return h
}

type resetPasswordHandler struct {
	registerHandler
}

func (h *resetPasswordHandler) Reset(user *model.TUser, password string) error {
	l := zap.L().With(zap.String("func", "修改密码"), zap.String("OperateId", h.OperateId))
	l.Info("用户修改密码")
	l.Info("解密密码")
	pwd, e := h.decodePwd(password)
	if e != nil {
		return e
	}

	l.Info("校验密码格式")
	if !h.verifyPwd(pwd) {
		return errors.New("密码格式不正确")
	}
	l.Info("更新用户密码")
	if e := database.DB.Model(user).Updates(map[string]any{
		"password":            utils.HashString(pwd),
		"password_updated_at": time.Now()}).Error; e != nil {
		return fmt.Errorf("密码更新失败: %s", e.Error())
	}
	return nil
}
