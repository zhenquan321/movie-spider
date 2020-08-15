package api

import (
	"movie_spider/model"
	"movie_spider/utils/spider"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// The MovieAPI provides handlers for managing movies.
type HomeAPI struct {
	DB MovieDatabase
}

func (a *HomeAPI) Home(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	var (
		page            int64
		sort            string
		order           int
		start           int64
		activeYear      int
		selZiCategories int
		selCategories   int
	)
	page, _ = strconv.ParseInt(ctx.DefaultQuery("page", "0"), 10, 64)
	start = page * 20
	activeYear, _ = strconv.Atoi(ctx.DefaultQuery("activeYear", ""))
	sort = ctx.DefaultQuery("sort", "released")
	selZiCategories, _ = strconv.Atoi(ctx.DefaultQuery("selZiCategories", ""))
	selCategories, _ = strconv.Atoi(ctx.DefaultQuery("selCategories", ""))

	var limit int64 = 20

	order = 1
	if sort == "id" {
		sort = "_id"
	}
	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}

	categories := spider.CategoriesStr()

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
	moviesCount := a.DB.CountMovie(condition)

	//生成年份列表
	var yearList [15]int
	for i := 0; i < 15; i++ {
		yearList[i] = 2020 - i /* 设置元素为 i + 100 */
	}

	ctx.HTML(200, "movie.html", gin.H{
		"title":           "牛逼的电影全在这",
		"movies":          movies,
		"allCount":        moviesCount,
		"pageCount":       limit,
		"page":            page,
		"yearList":        yearList,
		"activeYear":      activeYear,
		"categories":      categories,
		"selCategories":   selCategories,
		"selZiCategories": selZiCategories,
	})
}
