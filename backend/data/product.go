// This package provides a bunch of data structs and necessary functions
package data

type BasicProduct struct {
	ProductID             int
	Title                 string
	Brand                 BasicBrand
	Category              BasicCategory
	IsIndividuallyTracked bool
	PricePerTime          float64
	TimeUnit              string
}
