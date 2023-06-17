package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"quebrada_api/migrations"
)

func InitDB(connectionString string) *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	migrations.Migrate(db)

	return db
}

//func CheckConnection(connectionString string) *gorm.DB {
//
//	db, err := gorm.Open(postgres.New(postgres.Config{
//		DSN: connectionString,
//	}), &gorm.Config{})
//
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	defer db.C()
//
//	migrations.Migrate(db)
//
//	return db
//}
