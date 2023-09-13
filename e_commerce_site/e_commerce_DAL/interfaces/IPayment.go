package interfaces

// import "project1/models"


type Ipayment interface{
	CreatePayment( float32, int,float32,string)(string, error)
	// GetCustomerById()(*models.Customer,error)
}