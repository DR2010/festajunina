// Package models is a orders for packages
// -------------------------------------
// .../festajuninav2/models/orders.go
// -------------------------------------
package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Order is to be
type Order struct {
	SystemID         bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	ID               string        // random ID for order, yet to define algorithm
	ClientName       string        // Client Name
	ClientID         string        // Client ID in case they logon
	Atendente        string        // Pessoa atendendo
	Date             string        // Order Date
	Time             string        // Order Time
	Status           string        // Open, Completed, Cancelled
	EventID          string        // Every order belongs to an Event (Events are linked to Organisation)
	TimeStartServing string        // Open, Completed, Cancelled
	TimePlaced       string        // Open, Completed, Cancelled
	TimeCompleted    string        // Open, Completed, Cancelled
	TimeCancelled    string        // Open, Completed, Cancelled
	EatMode          string        // EatIn, TakeAway, Delivery
	PickUpTime       string        // EatIn, TakeAway, Delivery
	TotalGeral       string        // Delivery phone number
	Items            []Item
}

// Item represents a single item of an order
type Item struct {
	ID         string // Sequential number of the item
	PratoName  string // Dish ID or unique name from "Dishes"
	Quantidade string // Individual price
	Price      string // Individual price
	Total      string // Total Price
	Tax        string // GST
}

// SearchCriteria is what the client wants
type SearchCriteria struct {
	ID                   string // random ID for order, yet to define algorithm
	ClientName           string // Client Name
	ClientID             string // Client ID in case they logon
	Date                 string // Order Date
	Time                 string // Order Time
	Status               string // Open, Completed, Cancelled
	EatMode              string // EatIn, TakeAway, Delivery
	DeliveryMode         string // Internal, UberEats,
	DeliveryFee          string // Delivery Fee
	DeliveryLocation     string // Address
	DeliveryContactPhone string // Delivery phone number
}
