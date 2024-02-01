package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	Name       string
	URL        string `valid:"required~URL is required, url~URL URL is invalid"`
	EmployeeID string `valid:"required~EmployeeID is required ,stringlength(10|10),matches(^[E][M]//d{10}$)"`
}
