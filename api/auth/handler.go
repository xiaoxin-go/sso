package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sso/database"
	"sso/libs"
	"sso/model"
	"sso/pkg/auth"
	"sso/pkg/handler"
	"sso/utils"
	"strconv"
	"time"
)

type sendEmailCodeRequest struct {
	OperateId string `json:"operate_id" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

// SendEmailCode 发送验证码
func SendEmailCode(ctx *gin.Context) {
	params := sendEmailCodeRequest{}
	if e := ctx.ShouldBindJSON(&params); e != nil {
		libs.HttpParamsError(ctx, e.Error())
		return
	}
	rh := handler.NewRegisterHandler(params.OperateId)
	if e := rh.SendEmailCode(params.Email); e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	libs.HttpSuccess(ctx, nil, "发送成功")
}

type verifyEmailRequest struct {
	OperateId string `json:"operate_id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

// VerifyEmailCode 校验邮箱
func VerifyEmailCode(ctx *gin.Context) {
	params := verifyEmailRequest{}
	if e := ctx.ShouldBindJSON(&params); e != nil {
		libs.HttpParamsError(ctx, e.Error())
		return
	}
	rh := handler.NewRegisterHandler(params.OperateId)
	if e := rh.VerifyEmailCode(params.Email, params.Code); e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	libs.HttpSuccess(ctx, nil, "验证成功")
}

// GetOtpQrCode 获取otp二维码
func GetOtpQrCode(ctx *gin.Context) {
	operateId := ctx.Query("operate_id")
	if operateId == "" {
		libs.HttpParamsError(ctx, "operate_id不能为空")
		return
	}
	rh := handler.NewRegisterHandler(operateId)
	qrCode, e := rh.GetOtpQrCode()
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	// 返回图片
	ctx.Writer.Header().Set("Content-Type", "image/png")
	ctx.Writer.Header().Set("Content-Length", strconv.Itoa(len(qrCode)))
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Pragma", "no-cache")
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Writer.Header().Set("Content-Disposition", "inline; filename=qr.png")
	ctx.Writer.Header().Set("Content-Transfer-Encoding", "binary")

	// gin返回图片
	ctx.Data(200, "image/png", qrCode)
}

type verifyOtpCodeRequest struct {
	OperateId string `json:"operate_id" binding:"required"`
	Code      uint32 `json:"code" binding:"required"`
}

// VerifyOtpCode 校验otp
func VerifyOtpCode(ctx *gin.Context) {
	params := verifyOtpCodeRequest{}
	if e := ctx.ShouldBindJSON(&params); e != nil {
		libs.HttpParamsError(ctx, e.Error())
		return
	}
	rh := handler.NewRegisterHandler(params.OperateId)
	if e := rh.VerifyOtpCode(params.Code); e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	libs.HttpSuccess(ctx, nil, "验证成功")
}

type registerRequest struct {
	OperateId string `json:"operate_id" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	NameCn    string `json:"name_cn" binding:"required"`
}

// Register 注册
func Register(ctx *gin.Context) {
	params := registerRequest{}
	if e := ctx.ShouldBindJSON(&params); e != nil {
		libs.HttpParamsError(ctx, e.Error())
		return
	}
	user := model.TUser{
		Username: params.Username,
		Password: params.Password,
		NameCn:   params.NameCn,
	}
	rh := handler.NewRegisterHandler(params.OperateId)
	if e := rh.Register(&user); e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	libs.HttpSuccess(ctx, nil, "注册成功")
}

type loginRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	OperateId string `json:"operate_id" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

// Login 登录
func Login(ctx *gin.Context) {
	params := loginRequest{}
	if e := ctx.ShouldBindJSON(&params); e != nil {
		libs.HttpParamsError(ctx, e.Error())
		return
	}
	h := handler.NewLoginHandler(params.OperateId)
	user, sessionId, err := h.Login(params.Username, params.Password, params.Code)
	if err != nil {
		libs.HttpServerError(ctx, err.Error())
		return
	}
	ctx.SetCookie("sso_session_id", sessionId, 3600*24, "/", "", false, true)
	libs.HttpSuccess(ctx, user, "登录成功")
}

type retrievePasswordRequest struct {
	OperateId string `json:"operate_id" binding:"required"` // 操作ID
	Password  string `json:"password" binding:"required"`   // 加密后的密码
}

// RetrieveSendEmailCode 找回密码发送邮箱验证码
func RetrieveSendEmailCode(ctx *gin.Context) {
	params := sendEmailCodeRequest{}
	if e := ctx.ShouldBindJSON(&params); e != nil {
		libs.HttpParamsError(ctx, "参数解析异常: %s", e.Error())
		return
	}
	rh := handler.NewRetrieveHandler(params.OperateId)
	if e := rh.SendEmailCode(params.Email); e != nil {
		libs.ServerError(e.Error())
		return
	}
	libs.HttpSuccess(ctx, nil, "验证码发送成功")
}

// RetrievePassword 找回密码
func RetrievePassword(ctx *gin.Context) {
	params := retrievePasswordRequest{}
	if e := ctx.ShouldBindJSON(&params); e != nil {
		libs.HttpParamsError(ctx, "参数解析异常: %s", e.Error())
		return
	}
	rh := handler.NewRetrieveHandler(params.OperateId)
	if e := rh.Retrieve(params.Password); e != nil {
		libs.ServerError(e.Error())
		return
	}
	libs.HttpSuccess(ctx, nil, "重置成功")
}

// ResetPassword 修改密码
func ResetPassword(ctx *gin.Context) {
	params := retrievePasswordRequest{}
	if e := ctx.ShouldBindJSON(&params); e != nil {
		libs.HttpParamsError(ctx, e.Error())
		return
	}
	user, err := libs.GetUser(ctx)
	if err != nil {
		libs.HttpAuthorError(ctx, err.Error())
		return
	}
	h := handler.NewResetPasswordHandler(params.OperateId)
	if e := h.Reset(user, params.Password); e != nil {
		libs.HttpServerError(ctx, err.Error())
		return
	}
	libs.HttpSuccess(ctx, nil, "修改成功")
}

// GetPublicKey 获取公钥
func GetPublicKey(ctx *gin.Context) {
	operateId := ctx.Query("operate_id")
	if operateId == "" {
		libs.HttpParamsError(ctx, "operate_id不能为空")
		return
	}
	publicKey, privateKey, err := utils.GenerateKey()
	if err != nil {
		libs.HttpServerError(ctx, "公钥生成失败: %s", err.Error())
		return
	}
	// 保存私钥
	if e := database.R.HSet(operateId, "private_key", privateKey).Err(); e != nil {
		libs.HttpServerError(ctx, "保存私钥失败: %s", e.Error())
		return
	}
	database.R.Expire(operateId, 5*time.Minute)
	libs.HttpSuccess(ctx, publicKey, "ok")
}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx *gin.Context) {
	user, e := libs.GetUser(ctx)
	if e != nil {
		libs.AuthorError(e.Error())
		return
	}
	libs.HttpSuccess(ctx, user, "ok")
}

func GetMenus(ctx *gin.Context) {
	user, e := libs.GetUser(ctx)
	if e != nil {
		libs.HttpAuthorError(ctx, e.Error())
		return
	}
	zap.L().Debug(fmt.Sprintf("用户获取菜单------><%+v>", user))
	roleIds, err := new(model.TUserRole).PluckRoleIdsByUserId(user.Id)
	fmt.Println("role_ids-------->", roleIds)
	if err != nil {
		libs.HttpServerError(ctx, err.Error())
		return
	}
	result, err := auth.GetMenu(roleIds)
	if err != nil {
		msg := fmt.Sprintf("获取菜单信息异常: <%s>", err.Error())
		zap.L().Error(msg)
		libs.HttpServerError(ctx, msg)
		return
	}
	libs.HttpSuccess(ctx, result, "ok")
}

// SessionCheck 校验session是否有效
func SessionCheck(ctx *gin.Context) {
	sessionId := ctx.Query("session_id")
	appIdStr := ctx.Query("app_id")
	if sessionId == "" {
		libs.HttpParamsError(ctx, "session_id is required")
		return
	}
	if appIdStr == "" {
		libs.HttpParamsError(ctx, "app_id is required")
		return
	}
	appId, e := strconv.Atoi(appIdStr)
	if e != nil {
		libs.HttpParamsError(ctx, "app_id is invalid")
		return
	}
	// 从redis获取用户信息
	userId, e := database.R.HGet(sessionId, "user_id").Int()
	if e != nil {
		libs.HttpParamsError(ctx, "session_id is invalid")
		return
	}
	// 获取用户信息
	user := model.TUser{}
	if e := user.FirstById(userId); e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	// 获取平台信息
	platform := model.TPlatform{}
	if e := platform.FirstById(appId); e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	// 判断用户是否拥有平台权限
	// 获取用户下的租户
	userTenementIds, e := new(model.TTenementUser).PluckTenementIdsByUserId(userId)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	// 获取租户下的平台
	userPlatformIds, e := new(model.TTenementPlatform).PluckPlatformIdsByTenementIds(userTenementIds)
	if e != nil {
		libs.HttpServerError(ctx, e.Error())
		return
	}
	for _, v := range userPlatformIds {
		if v == appId {
			libs.HttpSuccess(ctx, user, "ok")
			return
		}
	}
	libs.HttpAuthorError(ctx, "用户没有访问平台的权限")
}
