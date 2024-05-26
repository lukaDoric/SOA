package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homeGetHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Welcome to Dojo!")
}
