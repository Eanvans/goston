package api

import (
	"gostonc/internal/app"
	"gostonc/internal/app/errcode"
	"gostonc/internal/router/req"
	"gostonc/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UserRegiser(c *gin.Context) {
	resp := app.NewResponse(c)
	data := req.RegisterUserReq{}
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
	data := req.UserPurchaseReq{}
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

func UserLogin(c *gin.Context) {
	param := req.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	_, err := service.DoLogin(param.Username, param.Password)
	if err != nil {
		logrus.Errorf("service.DoLogin err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	// token, refreshToken, tokenExpiredDuration, err := app.GenerateDoubleToken(user)
	// if err != nil {
	// 	logrus.Errorf("app.GenerateDoubleTokenerr: %v", err)
	// 	response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	// 	return
	// }

	response.ToResponse(gin.H{
		//"token":                  token, // access token
		// "token_expired_duration": tokenExpiredDuration,
		// "refresh_token":          refreshToken,
	})
}
