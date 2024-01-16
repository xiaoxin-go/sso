package log

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
	handler.NewResults = func() any {
		return make([]*model.TLog, 0)
	}
}
