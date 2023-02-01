package main

import (
	"github.com/gin-gonic/gin"
	"github.com/melrose12138/zhgg-dy/router"
	"github.com/melrose12138/zhgg-dy/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
