package entities

import "time"

type Category struct {
	Id          uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
