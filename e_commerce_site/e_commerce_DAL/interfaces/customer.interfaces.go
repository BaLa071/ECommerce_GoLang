package interfaces

import (
	"e_commerce_site/e_commerce_DAL/models"
)

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerDBResponse, error)
	CreateTokens(token *models.Token) (string, error)
	UpdatePassword(Password *models.UpdatePassword) (*models.CustomerDBResponse, error)
	UpdateCustomer(cus *models.UpdateRequest) (*models.CustomerDBResponse, error)
	DeleteCustomer(cus *models.DeleteRequest) error
	GetByCustomerId(res string) (*models.Customer, error)
	IsValidUser(cusomter *models.Customer) (bool)
}
