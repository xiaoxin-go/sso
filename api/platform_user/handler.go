package platform_user

import (
	"sso/libs"
	"sso/model"
)

type Handler struct {
	libs.Controller
}

var handler *Handler

func init() {
	handler = &Handler{}
	handler.NewInstance = func() libs.Instance {
		return new(model.TPlatformUser)
	}
	handler.NewResults = func() any {
		return &[]*model.TPlatformUser{}
	}
}
