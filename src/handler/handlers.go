package handler

import (
	"encoding/json"
	"github.com/bachelor-service/model"
	school "github.com/bachelor-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IHandler interface {
	GetSchoolsRespond(c *gin.Context)
	GetSchoolsByCountryRespond(c *gin.Context)
	CreateLocationRespond(c *gin.Context)
}

type Handler struct {
	SService school.IService
}

func (h Handler) GetSchoolsRespond(c *gin.Context) {
	var offset, _ = strconv.ParseInt(c.Query("offset"), 10, 64)
	var schools = make([]model.School, 0)

	h.SService.GetSchools(int(offset), &schools)

	c.IndentedJSON(http.StatusOK, schools)
}

func (h Handler) GetSchoolsByCountryRespond(c *gin.Context) {

	var schools []model.School
	var country = c.Param("name")

	h.SService.GetSchoolsByCountry(&schools, country)
	c.IndentedJSON(http.StatusOK, schools)
}

func (h Handler) CreateLocationRespond(c *gin.Context) {
	location := model.Location{}
	buf := make([]byte, 1024)
	c.Request.Body.Read(buf)
	json.Unmarshal(buf, &location)

	//h.SService.
}
