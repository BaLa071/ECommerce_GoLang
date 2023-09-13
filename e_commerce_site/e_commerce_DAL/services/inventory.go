package services

import (
	"context"
	"e_commerce_site/e_commerce_DAL/models"
	"e_commerce_site/e_commerce_DAL/interfaces"
	"fmt"


	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Inventory struct {
	ctx             context.Context
	itemCollection *mongo.Collection
	inventoryCollection *mongo.Collection
}

func InitInventory(itemCollection *mongo.Collection, inventoryCollection  *mongo.Collection, ctx context.Context) interfaces.Inventory {
	return &Inventory{ctx, itemCollection,inventoryCollection}
}



func (i *Inventory) CreateInventory(in []*models.Inventory) (*mongo.InsertManyResult, error) {
	// fmt.Println(in[0])
	// mcoll := config.GetCollection("inventory_SKU", "items")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{}
	result, err := i.itemCollection.Find(ctx, filter, options.Find())
	var inventory []*models.Inventory_SKU
	// fmt.Println(result)
	if err != nil {
		fmt.Println("error1")

		fmt.Println(err.Error())
		return nil, err
	}
	for result.Next(ctx) {
		post := &models.Inventory_SKU{}
		err := result.Decode(post)
		if err != nil {
			fmt.Println("Error decoding document:", err)
			// Print the problematic document or relevant information for debugging.
			fmt.Println("Document:", result.Current)
			return nil, err
		}
		inventory = append(inventory, post)
	}
	if err := result.Err(); err != nil {
		fmt.Println("error3")
		return nil, err
	}
	fmt.Println(inventory)
	n := 0
	for j := 0; j < len(in); j++ {
		for i := n; i < len(inventory); i++ {
			if in[j].Item == inventory[i].Options.Ruling{
			in[j].Skus = append(in[j].Skus, *inventory[i])
				}
		}
	
	}
	fmt.Println("in", in)
	inv := []interface{}{}
	for v := 0; v < len(in); v++ {
		inv = append(inv, in[v])
	}
	// inv := []interface{}(in)
	res, err := i.inventoryCollection.InsertMany(context.Background(), inv)
	if err != nil {
		fmt.Println("error4")
		return nil, err
	}
	return res, nil
}

func (i *Inventory) DeleteItems(item string, sku string, quantity float32) string {
	filter := bson.D{
		{Key: "item", Value: item},
		{Key: "skus.sku", Value: sku},
		{Key: "skus.quantity", Value: bson.D{{Key: "$gte", Value: quantity}}}, // Match the specific SKU within the "skus" array by SKU name.
	}

	update := bson.D{
		{Key:  "$inc", Value: bson.D{
			{Key: "skus.$.quantity", Value: -quantity}, // Decrement the "quantity" field by decrementAmount.
		}},
	}
	// fmt.Println("hello")
	res, err := i.inventoryCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
		return "failed"
	}
	fmt.Println(res)
	return "success"

}

func (i *Inventory) GetAllItems() ([]models.Inventory, error) {
	fmt.Println("GetAllItems2")

	filter := bson.D{} // An empty filter matches all documents.

	cursor, err := i.inventoryCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.Inventory

	for cursor.Next(context.Background()) {
		var result models.Inventory
		if err := cursor.Decode(&result); err != nil {

			return nil, err
		}
		results = append(results, result)

	}

	if err := cursor.Err(); err != nil {
		fmt.Println("err")

		return nil, err
	}

	return results, nil
}


func (i *Inventory) GetInventoryItemByItemName(itemName string) (*models.Inventory, error) {
	filter := bson.D{{Key: "item", Value: itemName}}
	fmt.Println("done")
	var result models.Inventory
	err := i.inventoryCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Println("error in decoding")
		return nil, err
	
	}
	fmt.Println(result.Features)
	
	return &result, nil
}
