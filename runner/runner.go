package runner

import (
	"fmt"
	"log"
	"net/http"

	"movie_spider/config"

	"github.com/gin-gonic/gin"
)

// Run starts the http server
func Run(engine *gin.Engine, conf *config.Configuration) {
	var httpHandler http.Handler = engine

	addr := fmt.Sprintf("%s:%d", conf.Server.ListenAddr, conf.Server.Port)
	fmt.Println("Started Listening for plain HTTP connection on " + addr)
	log.Fatal(http.ListenAndServe(addr, httpHandler))
}
