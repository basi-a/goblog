package main

import (
	"github.com/gin-gonic/autotls"
	"goblog/global"
)

func main() {
	db, _ := global.DB.DB()
	defer db.Close()
	serverStart()
}

func serverStart()  {
	r := global.RouterEngine
	if global.Config.TLS {
		autotls.Run(r, global.Config.Domains...)
	}
}