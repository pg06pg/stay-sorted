package migration

import (
	"github.com/pg06pg/stay-sorted/models/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		entities.User{},
		entities.House{},
		entities.HousePicture{},
	)
}
