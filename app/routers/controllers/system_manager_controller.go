package controllers

import (
	"github.com/gin-gonic/gin"
)

type SystemManagerController struct{}

func (SystemManagerController) Index(c *gin.Context) {
	c.String(200, "ok")
}
