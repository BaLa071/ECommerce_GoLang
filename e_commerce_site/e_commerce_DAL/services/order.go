package services

import (
	"context"
	"e_commerce_site/e_commerce_DAL/interfaces"
	"e_commerce_site/e_commerce_DAL/models"
	"fmt"

	// "strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService struct {
	client              *mongo.Client
	OrderCollection     *mongo.Collection
	InventoryCollection *mongo.Collection
	ctx                 context.Context
}

var Amt float32

func InitOrderService(client *mongo.Client, collection1 *mongo.Collection, collection2 *mongo.Collection, ctx context.Context) interfaces.IOrder {
	return &OrderService{client, collection1, collection2, ctx}
}

func (p *OrderService) CreateOrder(input *models.Orders) (models.Orders, error) {
	var basePrice float32
	var discountvalue float32
	var total float32
	for i, val := range input.Items {
		filter := bson.M{"sku": val.Sku}
		inventoryResult := p.InventoryCollection.FindOne(p.ctx, filter)
		var inventoryDocument bson.M
		if err := inventoryResult.Decode(&inventoryDocument); err != nil {
			return models.Orders{}, err
		}
		price := inventoryDocument["price"].(bson.M)
		quantity := inventoryDocument["quantity"].(float64)
		if quantity < float64(val.Quantity) {
			fmt.Println("Lack of Quantity")
			return models.Orders{}, nil
		}
		price64 := price["base"].(float64)
		discount := price["discount"].(float64)
		basePrice = float32(price64)
		discountvalue = float32(discount)
		fmt.Println(price64, discount, basePrice)
		input.Items[i].Price = val.Quantity * basePrice
		input.Items[i].Discount = discountvalue
		input.Items[i].PreTaxTotal = input.Items[i].Price - ((discountvalue / 100) * 100)
		total = total + input.Items[i].PreTaxTotal
		input.Items[i].Total = input.Items[i].PreTaxTotal
		fmt.Println(val.Price, val.Discount)
	}
	Amt = total
	input.TotalAmount = total
	insertResult, err1 := p.OrderCollection.InsertOne(p.ctx, &input)
	if err1 != nil {
		return models.Orders{}, err1
	}
	insertedID := insertResult.InsertedID.(primitive.ObjectID)

	filter := bson.M{"_id": insertedID}
	var result models.Orders
	err2 := p.OrderCollection.FindOne(p.ctx, filter).Decode(&result)
	if err2 != nil {
		return result, err2
	}
	fmt.Println("Success")
	return result, nil
}

func (p *OrderService) RemoveOrder(Customer_ID string) (string, error) {
	filter := bson.M{"customerid": Customer_ID}
	_, err := p.OrderCollection.DeleteOne(p.ctx, filter)
	if err != nil {
		return "Unable to delete", err
	}
	return "Deleted Successfully", nil
}

func (p *OrderService) GetAllOrder(CustomerId string) (*models.Orders, error) {
	filter := bson.M{"customerid": CustomerId}
	var res *models.Orders
	result := p.OrderCollection.FindOne(p.ctx, filter)
	err := result.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *OrderService) UpdateOrder(input *models.UpdateDetailsModel) (string, error) {
    var basePrice float32
    var discountvalue float32
    var total float32
    filter := bson.M{"customerid": input.Customer_ID}
    inventoryResult := p.InventoryCollection.FindOne(p.ctx, filter)
    var inventoryDocument bson.M
    if err := inventoryResult.Decode(&inventoryDocument); err != nil {
        return "Empty", err
    }
    price := inventoryDocument["price"].(bson.M)
    quantity := inventoryDocument["quantity"].(float64)
    if quantity < float64(input.Quantity) {
        fmt.Println("Lack of Quantity")
        return "Empty", nil
    }
    price64 := price["base"].(float64)
    discount := price["discount"].(float64)
    discountvalue = float32(discount)
    basePrice = float32(price64)
    fmt.Println("The Price is ", basePrice)
    input.Price = input.Quantity * basePrice
    input.Discount = discountvalue
    input.PreTaxTotal = input.Price - ((discountvalue / 100) * 100)
    total = input.PreTaxTotal
    input.TotalAmount = total
    input.Total = input.PreTaxTotal

    // Define a map for totalamount to ensure it's structured as an embedded document
    totalAmount := bson.M{"TotalAmount": input.TotalAmount}

    update := bson.M{
        "$set": bson.M{
            "totalamount": totalAmount, // Update totalamount field with the embedded document
        },
        "$push": bson.M{
            "items": bson.M{
                "sku":         input.Sku,
                "quantity":    input.Quantity,
                "price":       input.Price,
                "discount":    input.Discount,
                "pretaxtotal": input.PreTaxTotal,
                "total":       input.Total,
            },
        },
    }
    _, err1 := p.OrderCollection.UpdateOne(context.Background(), filter, update)
    if err1 != nil {
        return "updation failed", err1
    }
    return "Updation Success", nil
}



func (p *OrderService) AddOrder(input *models.UpdateDetailsModel) (string, error) {
	var store models.Orders
	var basePrice float32
	var discountvalue float32
	var total float32
	filter := bson.M{"sku": input.Sku}
	inventoryResult := p.InventoryCollection.FindOne(p.ctx, filter)
	var inventoryDocument bson.M
	if err := inventoryResult.Decode(&inventoryDocument); err != nil {
		return "Empty", err
	}
	price := inventoryDocument["price"].(bson.M)
	quantity := inventoryDocument["quantity"].(float64)
	if quantity < float64(input.Quantity) {
		fmt.Println("Lack of Quantity")
		return "Empty", nil
	}
	price64 := price["base"].(float64)
	discount := price["discount"].(float64)
	discountvalue = float32(discount)
	basePrice = float32(price64)
	fmt.Println("The Price is ", basePrice)
	fmt.Println("The quantity is ", input.Quantity)
	fmt.Println("The two values are ", input.Quantity, " ", basePrice)
	input.Price = input.Quantity * basePrice
	input.Discount = discountvalue
	input.PreTaxTotal = input.Price - ((discountvalue / 100) * 100)
	fmt.Println("Total amount ", input.TotalAmount)
	filter1 := bson.M{"customerid": input.Customer_ID}
	err3 := p.OrderCollection.FindOne(context.Background(), filter1).Decode(&store)
	if err3 != nil {
		return "updation failed", err3
	}
	total = input.PreTaxTotal + store.TotalAmount
	input.TotalAmount = total
	fmt.Println("Total amount ", input.TotalAmount)
	input.Total = input.PreTaxTotal
	fmt.Println("Decoded total amount", store.TotalAmount)
	itemsToAdd := []models.Items{
		{
			Sku:         input.Sku,
			Quantity:    input.Quantity,
			Price:       input.Price,
			Discount:    input.Discount,
			PreTaxTotal: input.PreTaxTotal + store.TotalAmount,
			Total:       input.Total,
		},
	}

	find := bson.M{
		"$set": bson.M{
			"totalamount": input.TotalAmount,
		},
	}
	update := bson.M{
		"$push": bson.M{
			"items": bson.M{
				"$each": itemsToAdd,
			},
		},
	}
	_, err2 := p.OrderCollection.UpdateOne(context.Background(), filter1, find)
	if err2 != nil {
		return "updation failed", err2
	}
	_, err1 := p.OrderCollection.UpdateOne(context.Background(), filter1, update)
	if err1 != nil {
		return "updation failed", err1
	}
	return "Updation Success", nil

}
