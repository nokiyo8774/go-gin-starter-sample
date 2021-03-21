package server

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/foods")
	{
		ctrl := controllers.FoodController{}
		u.GET("", ctrl.Index)
	}

	return r
}
