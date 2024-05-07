package global

import (
	"log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var RouterEngine *gin.Engine

func InitRouter() {

	setMode(Config.Mode)
	r := gin.Default()
	if err := r.SetTrustedProxies(Config.TrustedProxies); err != nil {
		log.Fatalln(err)
	}
	if Config.Mode == "debug" {
		// default is "debug/pprof"
		pprof.Register(r)
	}

	RouterEngine = r
}

func setMode(mode string) {

	//设置运行模式
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		log.Fatalln("Your run mode is set to", mode, ". Must be debug or release!!!!")
	}
}
