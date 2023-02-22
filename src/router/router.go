package router

import (
	"github.com/mateusVedoy/go-first-rest-api/gin-gonic/src/controller"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/albums", controller.GetAlbums)

	router.Run("localhost:2222")
}