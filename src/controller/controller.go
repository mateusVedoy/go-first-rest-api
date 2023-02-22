package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mateusVedoy/go-first-rest-api/src/model"
)

var albums = []models.Album{
	{
		Id:    "12321",
		Title: "First Album",
		Gender: models.Gender{
			Id:   "2352",
			Name: "Gender One",
		},
		Category: models.Category{
			Id:   "3465",
			Name: "Category One",
		},
	},
	{
		Id:    "235235",
		Title: "Second Album",
		Gender: models.Gender{
			Id:   "2352",
			Name: "Gender One",
		},
		Category: models.Category{
			Id:   "3465",
			Name: "Category One",
		},
	},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		albums,
	)
}
