package repository

import (
	"norht-gowind/src/common/database"
	model "norht-gowind/src/modules/employees/models"
)

func GetAllEmployees() ([]model.Employee, error) {
	var employees []model.Employee
	result := database.DB.Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}
	return employees, nil
}
