package models

import "gorm.io/gorm"

const ApplyToDo = 0
const ApplyRefuse = 1
const ApplyDONE = 2

type FriendApply struct {
	gorm.Model
	ID     int
	State  int  `gorm:"notnull:default:0"`
	FromID int  `gorm:"notnull;uniqueIndex:idx_from_to" json:"from_id"`
	ToID   int  `gorm:"notnull;uniqueIndex:idx_from_to" json:"to_id"`
	From   User `gorm:"foreignKey:FromID"`
	TO     User `gorm:"foreignKey:ToID"`
}

type ApiFriendApply struct {
	gorm.Model
	ID     int
	State  int
	FromID int
	ToID   int
}

// CreateFriendApply create a friend apply
func CreateFriendApply(db *gorm.DB, apply *FriendApply) (err error) {
	err = db.Create(apply).Error
	if err != nil {
		return err
	}
	return nil
}
