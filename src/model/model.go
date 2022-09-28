package model

type School struct {
	// gorm.Model
	ID           int64  `gorm:"primaryKey"`
	Name         string `json:"name"`
	Location_ref int
	Location     Location `gorm:"foreignKey:location_ref"`
}

type Location struct {
	// gorm.Model
	ID      int64 `gorm:"primaryKey"`
	Country string
	City    string
	Schools []School `gorm:"foreignKey:location_ref"`
}
