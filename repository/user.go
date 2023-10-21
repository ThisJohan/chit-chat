package repository

import "gorm.io/gorm"

type UserRepo struct {
	*gorm.DB
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (r *UserRepo) CreateUser(user *User) *gorm.DB {
	return r.DB.Create(user)
}

func (r *UserRepo) FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return r.DB.Model(&User{}).Take(dest, conds...)
}

func (r *UserRepo) FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return r.FindUser(dest, "email = ?", email)
}
