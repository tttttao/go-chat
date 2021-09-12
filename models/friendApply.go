package models

import "gorm.io/gorm"

const ApplyToDo = 0
const ApplyRefuse = 1
const ApplyDONE = 2

type FriendApply struct {
	gorm.Model
	ID     int
	State  int  `gorm:"notnull:default:0"`
	FromID int  `gorm:"notnull"`
	ToID   int  `gorm:"notnull"`
	From   User `gorm:"foreignKey:FromID"`
	TO     User `gorm:"foreignKey:ToID"`
}
