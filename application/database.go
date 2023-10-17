package application

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func (a *App) connectDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	a.db = db
}

func migrate(tables ...interface{}) error {
	return db.AutoMigrate(tables...)
}
