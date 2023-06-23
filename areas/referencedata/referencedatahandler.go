// Package referencedatahandler Handler for dishes web
// -----------------------------------------------------------
// .../src/restauranteweb/areas/disherhandler/disheshandler.go
// -----------------------------------------------------------
package referencedatahandler

import (
	"context"
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// This is the template to display as part of the html template
//

// RestEnvVariables = restaurante environment variables
//
type ReferenceDataCache struct {
	APIMongoDBLocation    string // location of the database localhost, something.com, etc
	APIMongoDBDatabase    string // database name
	APIAPIServerPort      string // collection name
	APIAPIServerIPAddress string // apiserver name
	WEBDebug              string // debug
	RecordCurrencyTick    string // debug
	RunningFromServer     string // debug
	WEBServerPort         string // collection name
	ConfigFileFound       string // collection name
	ApplicationID         string // collection name
	UserID                string // collection name
	AppFestaJuninaEnabled string
	AppBelnorthEnabled    string
	AppBitcoinEnabled     string
}

/*
CREATE TABLE accounts (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
        last_login TIMESTAMP

	ID                    VARCHAR ( 255 ) UNIQUE NOT NULL, // location of the database localhost, something.com, etc
	APIMongoDBLocation    VARCHAR ( 255 ), // location of the database localhost, something.com, etc
	APIMongoDBDatabase    VARCHAR ( 255 ),  // database name
	APIAPIServerPort      VARCHAR ( 255 ),  // collection name
	APIAPIServerIPAddress VARCHAR ( 255 ), // apiserver name
	WEBDebug              string // debug
	RecordCurrencyTick    string // debug
	RunningFromServer     string // debug
	WEBServerPort         string // collection name
	ConfigFileFound       string // collection name
	ApplicationID         string // collection name
	UserID                string // collection name
	AppFestaJuninaEnabled string
	AppBelnorthEnabled    string
	AppBitcoinEnabled     string



);
*/

// Insert saves one book to the database
func (refdatacachex *ReferenceDataCache) Insert(refdatacache ReferenceDataCache) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into books (title, author_id, publication_year, slug, description, created_at, updated_at)
            values ($1, $2, $3, $4, $5) returning id`

	var newID int
	err := db.QueryRowContext(ctx, stmt,
		refdatacache.APIMongoDBLocation,
		refdatacache.APIMongoDBDatabase,
		refdatacache.APIAPIServerPort,
		refdatacache.APIAPIServerIPAddress,
		refdatacache.WEBDebug,
		refdatacache.RecordCurrencyTick,
		refdatacache.RunningFromServer,
		refdatacache.WEBServerPort,
		refdatacache.ConfigFileFound,
		refdatacache.ApplicationID,
		refdatacache.UserID,
		refdatacache.AppFestaJuninaEnabled,
		refdatacache.AppBelnorthEnabled,
		refdatacache.AppBitcoinEnabled,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}
