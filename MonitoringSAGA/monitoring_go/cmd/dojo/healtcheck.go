package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func livenessHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Live!")
}

func readinessHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ready!")
}
