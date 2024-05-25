package service

import (
	model "norht-gowind/src/modules/customers/models"
	repository "norht-gowind/src/modules/customers/repositories"
)

func GetAllCustomers() ([]model.Customer, error) {
	return repository.FindAllCustomers()
}
