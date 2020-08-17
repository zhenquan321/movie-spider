package api

import (
	"errors"
	"movie_spider/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// The MovieDatabase interface for encapsulating database access.
type MovieDatabase interface {
	GetMovieByIDs(ids []primitive.ObjectID) []*model.Movie
	GetMovieByID(id primitive.ObjectID) *model.Movie
	GetMovieByName(name string) *model.Movie
	DeleteMovieByID(id primitive.ObjectID) error
	FindOneAndReplace(movie *model.Movie) *mongo.SingleResult
	CreateMovie(movie *model.Movie) error
	GetMovies(paging *model.Paging) []*model.Movie
	CountMovie(condition interface{}) string
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
		page            int64
		sort            string
		order           int
		start           int64
		activeYear      int
		selZiCategories int
	)
	page, _ = strconv.ParseInt(ctx.DefaultQuery("page", "0"), 10, 64)
	start = page * 20
	activeYear, _ = strconv.Atoi(ctx.DefaultQuery("activeYear", ""))
	sort = ctx.DefaultQuery("sort", "released")
	selZiCategories, _ = strconv.Atoi(ctx.DefaultQuery("selZiCategories", ""))

	var limit int64 = 20

	order = 1
	if sort == "id" {
		sort = "_id"
	}
	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}

	condition := make(bson.M)
	if activeYear > 0 {
		condition["released"] = strconv.Itoa(activeYear)
	}
	if selZiCategories > 0 {
		condition["typeId"] = selZiCategories
	} else {
		condition["typeId"] = 6
	}

	movies := a.DB.GetMovies(
		&model.Paging{
			Skip:      &start,
			Limit:     &limit,
			SortKey:   sort,
			SortVal:   order,
			Condition: condition,
		})
	// moviesCount := a.DB.CountMovie(condition)
	// allMoviesCount := a.DB.CountMovie(nil)

	ctx.JSON(200, movies)
}
