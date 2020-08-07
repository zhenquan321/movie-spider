package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"

	"movie_spider/config"
	"movie_spider/database"
	"movie_spider/mode"
	"movie_spider/model"
	"movie_spider/router"
	"movie_spider/runner"
	"movie_spider/utils"
	"movie_spider/utils/spider"

	_ "movie_spider/statik"

	"github.com/spf13/viper"
)

var (
	// Version the version of TMA.
	Version = "unknown"
	// Commit the git commit hash of this version.
	Commit = "unknown"
	// BuildDate the date on which this binary was build.
	BuildDate = "unknown"
	// Mode the build mode
	Mode = mode.Dev
)

// 首次启动自动开启爬虫
func firstSpider() {

	hasHK := utils.RedisDB.Exists("detail_links:id:14").Val()
	log.Println("hasHK", hasHK)
	// 不存在首页的key 则认为是第一次启动
	if hasHK == 0 {
		spider.Create().Start()
	}
}

// 初始化配置文件
func init() {
	viper.SetConfigType("json") // 设置配置文件的类型

	if err := viper.ReadConfig(bytes.NewBuffer(config.AppJsonConfig)); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}

func main() {
	vInfo := &model.VersionInfo{Version: Version, Commit: Commit, BuildDate: BuildDate}
	mode.Set(Mode)

	fmt.Println("Starting TMA version", vInfo.Version+"@"+BuildDate)
	rand.Seed(time.Now().UnixNano())
	conf := config.Get()

	db, err := database.New(conf.Database.Connection, conf.Database.Dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 初始化 redis 连接
	utils.InitRedisDB()
	defer utils.CloseRedisDB()
	//第一次启动检查爬虫
	firstSpider()
	// 启动定时爬虫任务
	utils.TimingSpider(func() {
		spider.Create().Start()
		return
	})

	engine := router.Create(db, vInfo, conf)
	runner.Run(engine, conf)
}
