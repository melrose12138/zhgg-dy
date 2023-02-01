package main

import (
	"github.com/gin-gonic/gin"
	"github.com/melrose12138/zhgg-dy/config"
	"github.com/melrose12138/zhgg-dy/database/mysql"
	"github.com/melrose12138/zhgg-dy/router"
	"log"
)

func main() {
	//go service.RunMessageServer()

	r := gin.Default()

	config.InitEnv()
	mysql.InitDB()
	router.InitRouter(r)

	port := ":" + config.Cfg.Server.Port
	if err := r.Run(port); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		log.Println(err)
	}
}
