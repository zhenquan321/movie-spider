package config

import (
	"github.com/lotteryjs/configor"
)

// Configuration is stuff that can be configured externally per env variables or config file (config.yml).
type Configuration struct {
	Server struct {
		ListenAddr      string `default:""`
		Port            int    `default:"80"`
		ResponseHeaders map[string]string
	}
	Database struct {
		Dbname     string `default:""`
		Connection string `default:""`
	}
}

// Get returns the configuration extracted from env variables or config file.
func Get() *Configuration {
	conf := new(Configuration)
	err := configor.New(&configor.Config{EnvironmentPrefix: "TenMinutesApi"}).Load(conf, "config.yml")
	if err != nil {
		panic(err)
	}
	return conf
}

/**

=========爬虫==========

参数说明：
app.spider_path: 爬虫路由
app.spider_path_name: 爬虫路由名称
app.debug_path: debug的路由
app.debug_path_name: debug的路由名称
cron.timing_spider: 定时爬虫的CRON表达式
ding.access_token: 钉钉机器人token
*/
var AppJsonConfig = []byte(`
{
  "app": {
    "port": ":8899",
    "spider_path": "/movies-spider",
    "spider_path_name": "MoviesSpider",
    "debug_path": "/debug",
    "debug_path_name": "Debug",
    "spider_mod": "api"
  },
  "redis": {
    "port": "6379",
    "addr": "localhost",
    "password": "",
    "db": 10
  },
  "cron": {
     "timing_spider": "0 0 1 * * ?"
  },
  "ding": {
     "access_token": ""
  }
}
`)
