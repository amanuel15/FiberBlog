package models

import "time"

type Blog struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	Title      string      `json:"title" gorm:"not null" validate:"required,min=1,max=100"`
	Body       string      `json:"body" gorm:"not null" validate:"required,min=10"`
	Category   string      `json:"category,omitempty"`
	Tags       []string    `json:"tags" gorm:"type:text"`
	AuthorID   uint        `json:"author_id,omitempty" gorm:"not null" validate:"omitempty"`
	Author     *User       `json:"author,omitempty" gorm:"foreignKey:AuthorID" validate:"omitempty"`
	Likes      []*User     `json:"likes,omitempty" gorm:"many2many:blog_likes;" validate:"omitempty"`
	References []Reference `json:"references,omitempty" validate:"omitempty"`
	Comments   []Comment   `json:"comments,omitempty" validate:"omitempty"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Reference struct {
	ID          string `json:"id,omitempty" gorm:"primaryKey"`
	Explanation string `json:"explanation,omitempty" gorm:"not null" validate:"required,min=3,max=100"`
	Link        string `json:"link" gorm:"not null" validate:"required,url"`
	BlogID      uint   `json:"blog_id,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
