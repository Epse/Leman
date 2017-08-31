// Simple db setup functions. DRY yuy
package dbase

import (
	"database/sql"
	"fmt"
	"github.com/Epse/Leman/backend/config"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// This just grabs all info from the config struct and returns the database infostring.
func GetInfoString(conf *config.BasicConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", conf.Database.User, conf.Database.Password, conf.Database.Database, conf.Database.URL)
}

// This creates the db object for you.
// Remember to do `defer db.Close()`!
func GetDBObject(infoString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", infoString)
	if err != nil {
		return db, errors.Wrap(err, "Could not open database connection")
	}
	return db, nil
}
