package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	v1.POST("/signup" , controllers.SignUpHandler)
	v1.POST("/login" , controllers.LoginHandler)
	v1.Use(middleware.JWTAuthMiddleware())
	{
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"404",
		})
	})
	return r
}

func RegisterValidator(){
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		 v.RegisterValidation("CheckMobileLayout", controllers.CheckMobileLayout)
	}
}