package data

import (
	"fmt"
	"testing"
)

func TestGetProductResponse(t *testing.T) {
	var product BasicProduct
	product.IsIndividuallyTracked = false
	product.PricePerTime = 2
	product.ProductID = 123
	product.TimeUnit = "day"
	product.Title = "Test"

	var brand BasicBrand
	brand.Title = "Brand"
	brand.BrandID = 456
	product.Brand = brand

	var cat BasicCategory
	cat.Title = "Categ"
	cat.CategoryID = 789
	product.Category = cat

	_, err := product.GetResponse()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestProductListResponse(t *testing.T) {
	var product BasicProduct
	product.IsIndividuallyTracked = false
	product.PricePerTime = 2
	product.ProductID = 123
	product.TimeUnit = "day"
	product.Title = "Test"

	var brand BasicBrand
	brand.Title = "Brand"
	brand.BrandID = 456
	product.Brand = brand

	var cat BasicCategory
	cat.Title = "Categ"
	cat.CategoryID = 789
	product.Category = cat

	var productArray [3]BasicProduct
	productArray[0] = product
	product.ProductID = 246
	productArray[1] = product
	product.ProductID = 482
	productArray[2] = product

	_, err := GetProductListResponse(productArray[:])
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
