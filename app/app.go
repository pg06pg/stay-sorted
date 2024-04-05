package app

import (
	"log"

	"github.com/pg06pg/stay-sorted/database"
)

func Bootstrap() {
	log.Println("Started Application!")
	LoadEnv()
	db := database.DatabaseConnection()
	router := Router(db)
	router.Run(":5000")
}
