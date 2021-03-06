package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"study/controllers"
	"study/logger"
	"study/middleware"
)

func Setup()*gin.Engine{
	r := gin.Default()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"ok")
	})
	v1.POST("/sendSms",controllers.SendSms)
	v1.POST("/checkSmsCode",controllers.CheckSmsCodeValid)
	v1.POST("/signup" , controllers.SignUpHandler)
	v1.POST("/signin",controllers.SignInHandler)
	v1.GET("/product/list",controllers.ProductListHandler)
	v1.GET("/product/detail",controllers.ProductDetailHandler)
	v1.GET("/product/category",controllers.ProductCategoryHandler)
	v1.GET("/category",controllers.GetCategory)
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.POST("/add/cart",controllers.AddCart)
		v1.GET("/get/cart",controllers.GetCart)
		v1.POST("/del/cart",controllers.DelCart)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"404",
		})
	})
	return r
}
