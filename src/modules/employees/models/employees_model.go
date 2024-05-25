package model

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	LastName        string
	FirstName       string
	Title           string
	TitleOfCourtesy string
	BirthDate       time.Time
	HireDate        time.Time
	Address         string
	City            string
	Region          string
	PostalCode      string
	Country         string
	HomePhone       string
	Extension       string
	Photo           []byte
	Notes           string
	ReportsTo       *int
	PhotoPath       string
}
