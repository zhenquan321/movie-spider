package spider

import (
	"movie_spider/utils"

	"github.com/spf13/viper"
)

// 定义 mod 的映射关系
var spiderModMap = map[string]utils.SpiderTask{
	"api":     &SpiderApi{},
	"WebPage": &utils.Spider{}}

func Create() utils.SpiderTask {

	mod := viper.GetString(`app.spider_mod`)

	return spiderModMap[mod]
}
