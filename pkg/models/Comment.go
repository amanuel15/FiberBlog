package models

import "time"

type Comment struct {
	ID              uint     `json:"id" gorm:"primaryKey"`
	Comment         string   `json:"comment" gorm:"not null"`
	WriterID        uint     `json:"writer_id,omitempty" gorm:"not null"`
	Writer          *User    `json:"writer,omitempty" gorm:"foreignKey:WriterID"`
	BlogID          uint     `json:"blog_id" gorm:"not null"`
	Blog            *Blog    `json:"blog,omitempty" gorm:"foreignKey:BlogID"`
	ReplayTo        *uint    `json:"replay_to,omitempty"`
	OriginalComment *Comment `json:"original_comment,omitempty" gorm:"foreignKey:ReplayTo"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
