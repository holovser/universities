package school

import (
	"github.com/bachelor-service/model"
	"gorm.io/gorm"
)

type IService interface {
	GetSchools(offset int, schools *[]model.School) []model.School
	CreateLocation(location model.Location)
	GetSchoolsByCountry(schools *[]model.School, country string) []model.School
}

type Service struct {
	Db *gorm.DB
}

func (s Service) GetSchools(offset int, schools *[]model.School) []model.School {
	s.Db.Limit(20).Offset(offset).Preload("Location").Find(schools)

	return *schools
	//db.Limit(20).Find(schools)
}

func (s Service) GetSchoolsByCountry(schools *[]model.School, country string) []model.School {

	var locations = make([]model.Location, 1)
	*schools = make([]model.School, 0)

	s.Db.Where("country = ?", country).Preload("Schools").Find(&locations)

	for _, location := range locations {
		*schools = append(*schools, location.Schools...)
	}

	return *schools

	//db.Where("").Limit(20).Offset(offset).Preload("Location").Find(schools)
	//db.Limit(20).Find(schools)
}

func (s Service) CreateLocation(location model.Location) {

}
