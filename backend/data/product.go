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

// Return an array of all products
func GetAllProducts(conf *config.BasicConfig) ([]BasicProduct, error) {
	var products []BasicProduct

	db, err := dbase.GetDBObject(dbase.GetInfoString(conf))
	if err != nil {
		return products, errors.Wrap(err, "Querying for all products")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Products")
	if err != nil {
		return products, errors.Wrap(err, "Was querying for all products")
	}

	i := 0
	for rows.Next() {
		var brandID int
		var categoryID int
		err := rows.Scan(&products[i].ProductID, &products[i].Title, &brandID, &categoryID, &products[i].IsIndividuallyTracked, &products[i].PricePerTime, &products[i].TimeUnit)

		if err != nil {
			return products, errors.Wrap(err, "Couldn't scan product.")
		}

		// Lets figure out the brand
		products[i].Brand, err = GetBrand(brandID, conf)
		if err != nil {
			return products, errors.Wrapf(err, "Was resolving brand with ID %d", brandID)
		}

		// And the category
		products[i].Category, err = GetCategory(categoryID, conf)
		if err != nil {
			return products, errors.Wrapf(err, "Was resolving category ID %d", categoryID)
		}

		i++
	}

	return products, nil
}
