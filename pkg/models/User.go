package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null" validate:"required,min=3,max=12"`
	Password  string    `json:"password" gorm:"not null"`
	Tags      []string  `json:"tags,omitempty" gorm:"type:text"`
	Likes     []*Blog   `json:"likes,omitempty" gorm:"many2many:liked_blogs;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	Name     string `json:"name" validate:"required,min=1,max=50"`
	Password string `json:"password" validate:"required,min=5,max=16"`
}

type DecodedUser struct {
	UserId uint `json:"userId"`
}
