package application

import (
	"github.com/ThisJohan/ChitChat/repository"
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

	updateMigrations()

	a.db = db
}

func updateMigrations() {
	migrate(&repository.User{})
}

func migrate(tables ...interface{}) error {
	return db.AutoMigrate(tables...)
}
