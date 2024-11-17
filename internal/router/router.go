package router

import (
	"gostonc/internal/router/api"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	// e.Use(cors.New(corsConfig))
	// {
	// 	// localOSS 路由
	// 	routeLocalOSS(e)
	// }

	// pprof api 性能分析用
	pprof.Register(e, "/monitor/pprof")

	r := e.Group("/v1")
	setVersionRoutes(r)

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})
	return e
}

func setVersionRoutes(r *gin.RouterGroup) {
	noAuthApi := r.Group("/")
	{
		noAuthApi.POST("/user/register", api.UserRegiser)
		noAuthApi.POST("/user/purchase", api.PurchaseTimespan)
	}
}
