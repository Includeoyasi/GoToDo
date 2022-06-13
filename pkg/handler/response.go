package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func NewErrorResponse(gctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	gctx.AbortWithStatusJSON(statusCode, error{Message: message})
}
