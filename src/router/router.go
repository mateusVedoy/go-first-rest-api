package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusVedoy/go-first-rest-api/src/config"
	"github.com/mateusVedoy/go-first-rest-api/src/controller"
)

func StartRouter() {

	PATH := config.ServeBASEURL()

	router := gin.Default()

	router.GET("/albums", controller.GetAlbums)

	router.Run(PATH)
}
