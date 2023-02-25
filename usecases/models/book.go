package models

import "time"

type Book struct {
	ID        uint
	Title     string
	Author    string
	Pages     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
