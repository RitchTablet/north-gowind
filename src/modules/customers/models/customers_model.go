package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerID   string `gorm:"type:varchar(5);primary_key"`
	CompanyName  string
	ContactName  string
	ContactTitle string
	Address      string
	City         string
	Region       string
	PostalCode   string
	Country      string
	Phone        string
	Fax          string
}
