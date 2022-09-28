package app

import (
	"github.com/bachelor-service/handler"
	"github.com/gin-gonic/gin"
)

type App struct {
	H handler.IHandler
}

func (a App) Start() {
	//dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//	Logger: logger.Default.LogMode(logger.Info),
	//})

	//if err != nil {
	//	panic("failed to connect database")
	//}

	router := gin.Default()
	a.configRoutes(router)
	router.Run("localhost:8080")
}

func (a App) configRoutes(router *gin.Engine) {
	router.GET("/schools", func(context *gin.Context) {
		a.H.GetSchoolsRespond(context)
	})

	router.GET("/country/:name/schools", func(context *gin.Context) {
		a.H.GetSchoolsByCountryRespond(context)
	})

	router.POST("/schools", func(context *gin.Context) {
		a.H.CreateLocationRespond(context)
	})
}
