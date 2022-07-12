package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jazz-service/model"
	"net/http"
)

// getAlbums responds with the list of all albums as JSON.
func getSchools(c *gin.Context) {

	// var location model.Location
	//var school model.School
	//db.Preload("Location").First(&school)

	var location model.Location
	location.School = make([]model.School, 0)

	db.Preload("School").First(&location)

	fmt.Println(location.School)

	c.IndentedJSON(http.StatusOK, location)
}
