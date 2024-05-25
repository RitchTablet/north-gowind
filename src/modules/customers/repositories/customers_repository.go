package repository

import (
	"norht-gowind/src/common/database"
	model "norht-gowind/src/modules/customers/models"
)

func FindAllCustomers() ([]model.Customer, error) {
	var customers []model.Customer
	result := database.DB.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}
