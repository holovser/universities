package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jazz-service/model"
	"net/http"
	"strconv"
)

// getAlbums responds with the list of all albums as JSON.
func getSchools(c *gin.Context) {

	var offset, _ = strconv.ParseInt(c.Param("offset"), 10, 64)

	// var location model.Location
	var schools []model.School
	db.Limit(20).Offset(int(offset)).Preload("Location").Find(&schools)

	//var location model.Location
	//location.School = make([]model.School, 0)

	//db.Preload("School").First(&location)

	//fmt.Println(location.School)

	c.IndentedJSON(http.StatusOK, schools)
}
