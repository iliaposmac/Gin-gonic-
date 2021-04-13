package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:""`
}

func newResponseError(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
