// This package provides a bunch of data structs and necessary functions
package data

import (
	"github.com/Epse/Leman/backend/config"
	"github.com/Epse/Leman/backend/dbase"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type BasicBrand struct {
	BrandID int
	Title   string
}

// When provided a brand ID and the lovely configuration struct
// returns either the brand as BasicBrand and nil,
// or an empty BasicBrand and an error
func GetBrand(brandID int, conf *config.BasicConfig) (BasicBrand, error) {
	var brand BasicBrand

	db, err := dbase.GetDBObject(dbase.GetInfoString(conf))
	if err != nil {
		return brand, errors.Wrapf(err, "Querying with brand id %d", brandID)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Brands WHERE BrandID=%d", brandID)
	if err != nil {
		return brand, errors.Wrapf(err, "Can't query for brand id %d", brandID)
	}

	// For reals, this should really only ever have one result.
	// It CAN only have one result, yet still
	for rows.Next() {
		err := rows.Scan(&brand.BrandID, &brand.Title)
		if err != nil {
			return brand, errors.Wrapf(err, "Can't read brand with id %d into struct", brandID)
		}
	}

	return brand, nil
}
