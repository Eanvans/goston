package api

import (
	"gostonc/internal/app"
	"gostonc/internal/app/errcode"
	"gostonc/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UserRegiser(c *gin.Context) {
	resp := app.NewResponse(c)
	data := RegisterUserReq{}
	if valid, errs := app.BindAndValid(c, &data); !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	u, err := service.UserRegister(data.Username, data.Password)
	if err != nil {
		resp.ToErrorResponse(errcode.ServerError)
		return
	}

	resp.ToResponse(u)
}

func PurchaseTimespan(c *gin.Context) {
	resp := app.NewResponse(c)
	data := UserPurchaseReq{}
	if valid, errs := app.BindAndValid(c, &data); !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		resp.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := service.UserPurchaseTimespan(data.UserID)
	if err != nil {
		resp.ToErrorResponse(err.(*errcode.Error))
		return
	}

	resp.ToResponse()
}
