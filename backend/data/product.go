// This package provides a bunch of data structs and necessary functions
package data

import (
	"encoding/json"
	"github.com/Epse/Leman/backend/config"
	"github.com/Epse/Leman/backend/dbase"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type BasicProduct struct {
	ProductID             int
	Title                 string
	Brand                 BasicBrand
	Category              BasicCategory
	IsIndividuallyTracked bool
	PricePerTime          float64
	TimeUnit              string
}

type ProductResponse struct {
	Product    BasicProduct
	Trackables []BasicTrackable
}

// When provided a product ID and our lovely configuration struct,
// returns either the product as BasicProduct and nil,
// or an empty BasicProduct and an error
func GetProduct(productID int, conf *config.BasicConfig) (BasicProduct, error) {
	var product BasicProduct

	db, err := dbase.GetDBObject(dbase.GetInfoString(conf))
	if err != nil {
		return product, errors.Wrapf(err, "Querying for product with id %d", productID)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Products WHERE ProductID=%d", productID)
	if err != nil {
		return product, errors.Wrapf(err, "Could not query for product with id %d", productID)
	}

	// This can really only ever have one result.
	var brandid int
	var categoryid int
	for rows.Next() {
		err := rows.Scan(&product.ProductID, &product.Title, &brandid, &categoryid, &product.IsIndividuallyTracked, &product.PricePerTime, &product.TimeUnit)
		if err != nil {
			return product, errors.Wrapf(err, "Could not read product with id %d into struct", productID)
		}
	}

	// Now we should look up the brand
	product.Brand, err = GetBrand(brandid, conf)
	if err != nil {
		return product, errors.Wrapf(err, "Was querying for product %d", productID)
	}

	// And the category...
	product.Category, err = GetCategory(categoryid, conf)
	if err != nil {
		return product, errors.Wrapf(err, "Was querying for product %d", productID)
	}

	return product, nil
}

func (product BasicProduct) GetResponse() ([]byte, error) {
	response, err := json.Marshal(product)
	if err != nil {
		return response, errors.Wrapf(err, "Was encoding product ID %d", product.ProductID)
	}
	return response, nil
}
