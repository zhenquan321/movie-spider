package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"movie_spider/api"
	"movie_spider/config"
	"movie_spider/database"
	"movie_spider/error"
	"movie_spider/model"
)

// Create creates the gin engine with all routes.
func Create(db *database.TenDatabase, vInfo *model.VersionInfo, conf *config.Configuration) *gin.Engine {
	g := gin.Default()
	g.Use(gin.Logger(), gin.Recovery(), error.Handler())
	g.NoRoute(error.NotFound())

	g.Use(func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		origin := ctx.Request.Header.Get("Origin")
		for header, value := range conf.Server.ResponseHeaders {
			if origin == "http://localhost:3000" && header == "Access-Control-Allow-Origin" {
				ctx.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			} else {
				ctx.Header(header, value)
			}
		}
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
	})

	userHandler := api.UserAPI{DB: db}
	postHandler := api.PostAPI{DB: db}
	movieHandler := api.MovieAPI{DB: db}
	homeHandler := api.HomeAPI{DB: db}

	postU := g.Group("/users")
	{
		postU.GET("", userHandler.GetUsers)
		postU.DELETE(":id", userHandler.DeleteUserByID)
	}

	postG := g.Group("/posts")
	{
		postG.GET("", postHandler.GetPosts)
		postG.POST("", postHandler.CreatePost)
		postG.GET(":id", postHandler.GetPostByID)
		postG.PUT(":id", postHandler.UpdatePostByID)
		postG.DELETE(":id", postHandler.DeletePostByID)
	}

	postM := g.Group("/movie")
	{
		postM.GET("", movieHandler.GetMovies)
		postM.GET(":id", movieHandler.GetMovieByIDs)
		postM.DELETE(":id", movieHandler.DeleteMovieByID)
	}

	g.LoadHTMLGlob("templates/**/*")
	g.Static("/static", "./static/")
	postH := g.Group("/home")
	{
		postH.GET("", homeHandler.Home)
	}

	g.GET("version", func(ctx *gin.Context) {
		ctx.JSON(200, vInfo)
	})

	return g
}
