package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusVedoy/go-first-rest-api/src/controller"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/albums", controller.GetAlbums)

	router.Run("localhost:2222")
}
