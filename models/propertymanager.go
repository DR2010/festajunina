// Package models is a property manager for packages
// -------------------------------------
// .../restauranteapi/models/dishes.go
// -------------------------------------
package models

import (
	"gopkg.in/mgo.v2/bson"
)

/*


System Owner/ People that would buy the system
	Strata Manager     (main user)
	Property Manager   (main user)

Client
	Landlord
	Tenant

People Involved
	Contractor
	Building Manager
	Finance Officer
	System Users

*/

// Dish is to be exported
type PropertyManager struct {
	SystemID bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Name     string        // name of the dish - this is the KEY, must be unique
	Type     string        // type of dish, includes drinks and deserts
	Price    string        // preco do prato multiplicar por 100 e nao ter digits

}
