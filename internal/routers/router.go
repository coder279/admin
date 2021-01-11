package routers

import (
	"github.com/gin-gonic/gin"
	"lianquan/global"
	"lianquan/internal/middleware"
	"lianquan/pkg/limiter"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine{
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeOut(global.AppSetting.DefaultContextTimeout))
	r.GET("/", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// callback 是 x
		// 将输出：x({\"foo\":\"bar\"})
		context.JSONP(http.StatusOK, data)
	})
	return r

}

