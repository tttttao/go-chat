package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int
	Name        string        `gorm:"unique;type:string;size:100;not null;default:''"`
	Email       string        `gorm:"unique;type:string;size:100;not null;default:''"`
	Password    []byte        `gorm:"not null;"`
	Avatar      string        `gorm:"not null;default:''"`
	Nickname    string        `gorm:"type:string;size:100;not null;default:''"`
	Friends     []*User       `gorm:"many2many:user_friends;"`
	AppliesFrom []FriendApply `gorm:"foreignKey:FromID;"`
	AppliesTo   []FriendApply `gorm:"foreignKey:ToID;"`
}

// CreateUser create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUsers get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUser get user by id
func GetUser(db *gorm.DB, User *User) (err error) {
	err = db.Where(User).Preload("AppliesFrom").First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

// DeleteUser delete user
func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
