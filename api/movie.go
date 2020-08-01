package api

import (
	"errors"
	"net/http"
	"strconv"

	"movie_spider/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// The MovieDatabase interface for encapsulating database access.
type MovieDatabase interface {
	GetMovieByIDs(ids []primitive.ObjectID) []*model.Movie
	GetMovieByName(name string) *model.Movie
	DeleteMovieByID(id primitive.ObjectID) error
	FindOneAndReplace(movie *model.Movie) *mongo.SingleResult
	CreateMovie(movie *model.Movie) error
	GetMovies(paging *model.Paging) []*model.Movie
	CountMovie() string
}

// The MovieAPI provides handlers for managing movies.
type MovieAPI struct {
	DB MovieDatabase
}

// GetMovieByIDs returns the movie by id
func (a *MovieAPI) GetMovieByIDs(ctx *gin.Context) {
	withIDs(ctx, "id", func(ids []primitive.ObjectID) {
		ctx.JSON(200, a.DB.GetMovieByIDs(ids))
	})
}

// DeleteMovieByID deletes the movie by id
func (a *MovieAPI) DeleteMovieByID(ctx *gin.Context) {
	withID(ctx, "id", func(id primitive.ObjectID) {
		if err := a.DB.DeleteMovieByID(id); err == nil {
			ctx.JSON(200, http.StatusOK)
		} else {
			if err != nil {
				ctx.AbortWithError(500, err)
			} else {
				ctx.AbortWithError(404, errors.New("movie does not exist"))
			}
		}
	})
}

// GetMovies returns all the movies
// _end=5&_order=DESC&_sort=id&_start=0 adapt react-admin
func (a *MovieAPI) GetMovies(ctx *gin.Context) {
	var (
		start int64
		end   int64
		sort  string
		order int
	)
	id := ctx.DefaultQuery("id", "")
	if id != "" {
		a.GetMovieByIDs(ctx)
		return
	}
	start, _ = strconv.ParseInt(ctx.DefaultQuery("_start", "0"), 10, 64)
	end, _ = strconv.ParseInt(ctx.DefaultQuery("_end", "10"), 10, 64)
	sort = ctx.DefaultQuery("_sort", "_id")
	order = 1

	if sort == "id" {
		sort = "_id"
	}

	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}

	limit := end - start
	movies := a.DB.GetMovies(
		&model.Paging{
			Skip:      &start,
			Limit:     &limit,
			SortKey:   sort,
			SortVal:   order,
			Condition: nil,
		})

	ctx.Header("X-Total-Count", a.DB.CountMovie())
	ctx.JSON(200, movies)
}
