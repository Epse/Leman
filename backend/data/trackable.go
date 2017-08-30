// This package provides a bunch of data structs and necessary functions
package data

type BasicTrackable struct {
	TrackableID        int
	Product            BasicProduct
	TotalQuantity      int
	LocationIIT        BasicLocation
	QuantitiyAvailable int
}
