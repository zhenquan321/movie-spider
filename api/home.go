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
		start      int64
		end        int64
		sort       string
		order      int
		activeYear int
	)

	start, _ = strconv.ParseInt(ctx.DefaultQuery("_start", "0"), 10, 64)
	end, _ = strconv.ParseInt(ctx.DefaultQuery("_end", "20"), 10, 64)
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

	limit := end - start
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
		"title":       "溜芒之道",
		"movies":      movies,
		"moviesCount": moviesCount,
		"yearList":    yearList,
		"activeYear":  activeYear,
	})
}
