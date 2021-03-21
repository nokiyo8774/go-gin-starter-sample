package controllers

import (
	"app/common/db"
	"app/models"
	"context"

	"github.com/gin-gonic/gin"
)

type FoodController struct{}

func (FoodController) Index(c *gin.Context) {
	res, err := models.Foods().All(context.Background(), db.DB)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, res)
}
