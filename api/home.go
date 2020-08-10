package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// The MovieAPI provides handlers for managing movies.
type HomeAPI struct {
	DB MovieDatabase
}

func (a *HomeAPI) Home(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")

	ctx.HTML(200, "movie.html", gin.H{
		"title": "hello gin " + strings.ToLower(ctx.Request.Method) + " method",
	})
}
