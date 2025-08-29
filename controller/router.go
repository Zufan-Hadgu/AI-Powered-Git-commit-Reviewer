package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(reviewcontroller *ReviewController) *gin.Engine{
	r := gin.Default()

	r.POST("/review",reviewcontroller.Review)

	return r

}