package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name   string
	Done   bool
	DoneOn time.Time
}
