package auth

import "github.com/gin-gonic/gin"

func Routers(e *gin.RouterGroup) {
	e.POST("/auth/send_email_code", SendEmailCode)                  // 发送邮箱验证码
	e.POST("/auth/verify_email_code", VerifyEmailCode)              // 验证邮箱
	e.GET("/auth/otp_qr_code", GetOtpQrCode)                        // 获取otp二维码
	e.POST("/auth/verify_otp_code", VerifyOtpCode)                  // 验证otp
	e.POST("/auth/register", Register)                              // 注册
	e.GET("/auth/public_key", GetPublicKey)                         // 获取公钥
	e.POST("/auth/login", Login)                                    // 登录
	e.POST("/auth/retrieve_send_email_code", RetrieveSendEmailCode) // 找回密码发送邮箱验证码
	e.POST("/auth/retrieve_password", RetrievePassword)             // 返回密码
	e.POST("/auth/reset_password", ResetPassword)                   // 重置密码
	e.GET("/auth/user_info", GetUserInfo)
	e.GET("/auth/session_check", SessionCheck) // 校验用户是否登录
	e.GET("/auth/menus", GetMenus)

}
