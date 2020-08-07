package api

import (
	"github.com/gin-gonic/gin"
)

// The MovieAPI provides handlers for managing movies.
type HomeAPI struct {
	DB MovieDatabase
}

func (a *HomeAPI) Home(ctx *gin.Context) {
	ctx.Status(200)

	ctx.HTML(200, "index.html", "flysnow_org")
}
