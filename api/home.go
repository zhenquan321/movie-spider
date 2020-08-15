package api

import (
	"movie_spider/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// The MovieAPI provides handlers for managing movies.
type HomeAPI struct {
	DB MovieDatabase
}

func (a *HomeAPI) Home(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	var (
		page       int64
		sort       string
		order      int
		start      int64
		activeYear int
	)

	page, _ = strconv.ParseInt(ctx.DefaultQuery("_page", "0"), 10, 64)
	start = page * 20
	activeYear, _ = strconv.Atoi(ctx.DefaultQuery("_activeYear", "0"))
	sort = ctx.DefaultQuery("_sort", "released")
	order = 1

	if sort == "id" {
		sort = "_id"
	}

	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}
	//生成年份列表
	var yearList [15]int
	for i := 0; i < 15; i++ {
		yearList[i] = 2020 - i /* 设置元素为 i + 100 */
	}
	var limit int64 = 20
	movies := a.DB.GetMovies(
		&model.Paging{
			Skip:      &start,
			Limit:     &limit,
			SortKey:   sort,
			SortVal:   order,
			Condition: nil,
		})
	moviesCount := a.DB.CountMovie()

	ctx.HTML(200, "movie.html", gin.H{
		"title":      "牛逼的电影全在这",
		"movies":     movies,
		"allCount":   moviesCount,
		"pageCount":  limit,
		"page":       page,
		"yearList":   yearList,
		"activeYear": activeYear,
	})
}
