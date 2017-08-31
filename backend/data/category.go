// This package provides a bunch of data structs and necessary functions
package data

import (
	"github.com/Epse/Leman/backend/config"
	"github.com/Epse/Leman/backend/dbase"
	"github.com/pkg/errors"
)

type BasicCategory struct {
	CategoryID int
	Title      string
}

func GetCategory(categoryID int, conf *config.BasicConfig) (BasicCategory, error) {
	var category BasicCategory

	db, err := dbase.GetDBObject(dbase.GetInfoString(conf))
	if err != nil {
		return category, errors.Wrapf(err, "Was querying for category %d", categoryID)
	}

	rows, err := db.Query("SELECT * FROM Categories WHERE CategoryID=%d", categoryID)
	if err != nil {
		return category, errors.Wrapf(err, "Was querying for category %d", categoryID)
	}

	// This CAN really ever only return one result, but yeah
	for rows.Next() {
		err := rows.Scan(&category.CategoryID, &category.Title)
		if err != nil {
			return category, errors.Wrapf(err, "Was querying for category %d", categoryID)
		}
	}

	return category, nil
}
