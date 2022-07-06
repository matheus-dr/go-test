package entities

import "time"

type Author struct {
	Id        uint
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
