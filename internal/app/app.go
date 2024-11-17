package app

import (
	"gostonc/internal/app/errcode"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalRows int64 `json:"total_rows"`
}

type ResponseData struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	TraceHost string      `json:"trace_host"`
	Data      interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	//hostname, _ := os.Hostname()
	if data == nil {
		data = gin.H{
			"code": 0,
			"msg":  "success",
		}
	} else {
		data = gin.H{
			"code": 0,
			"msg":  "success",
			"data": data,
		}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseArray(arr interface{}) {
	r.ToResponse(gin.H{
		"list": arr,
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}

type ValidError struct {
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		errs = append(errs, &ValidError{
			Message: err.Error(),
		})

		return false, errs
	}

	return true, nil
}
