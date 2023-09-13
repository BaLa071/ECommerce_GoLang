package interfaces

import "e_commerce_site/e_commerce_DAL/models"

type IOrder interface {
	CreateOrder(input *models.Orders) (models.Orders, error)
	RemoveOrder(Customer_ID string) (string, error)
	GetAllOrder(CustomerId string) (*models.Orders, error)
	UpdateOrder(input *models.UpdateDetailsModel)(string,error)
	AddOrder(input *models.UpdateDetailsModel) (string, error)
}
