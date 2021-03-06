package routers

import (
	"app/common/config/logging"
	"app/routers/controllers"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	logger := logging.Setup()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	u := r.Group("/system-manager")
	{
		ctrl := controllers.SystemManagerController{}
		u.GET("health-check", ctrl.Index)
	}

	u = r.Group("/foods")
	{
		ctrl := controllers.FoodController{}
		u.GET("", ctrl.Index)
	}

	return r
}
