package model

import "time"

type BookHistory struct {
	ID           int64     `gorm:"column:id;primary_key"  json:"id"`
	UserID       int64     `gorm:"column:user_id"  json:"user_id"`
	BookID       int64     `gorm:"column:book_id"  json:"book_id"`
	LastViewPage int64     `gorm:"column:last_view_page"  json:"last_view_page"`
	LastViewAt   time.Time `gorm:"column:last_view_at"  json:"last_view_at"`
}
