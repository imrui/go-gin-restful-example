package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-restful-example/conf"
	"go-gin-restful-example/models"
	"go-gin-restful-example/restful"
	"log"
)

func main() {
	err := models.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	restful.LoadRouters(r)
	log.Fatal(r.Run(conf.Cfg.Server.Addr))
}
