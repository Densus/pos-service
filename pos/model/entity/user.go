package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	UserName string `gorm:"type:varchar(64)" json:"user_name"`
	FullName string `gorm:"type:varchar(64)" json:"full_name"`
	Email string `gorm:"uniqueIndex;type:varchar(64)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Token string `gorm:"-" json:"-"`
	Role string `gorm:"type:varchar(10)" json:"role"`
}

func (User) TableName() string {
	return "users"
}
