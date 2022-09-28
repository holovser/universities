package di

import (
	"github.com/bachelor-service/handler"
	school "github.com/bachelor-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ProviderHandler(service school.IService) *handler.Handler {
	return &handler.Handler{SService: service}
}

func ProviderService(db *gorm.DB) *school.Service {
	return &school.Service{Db: db}
}

func ProviderDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
