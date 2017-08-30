// This package provides a bunch of data structs and necessary functions
package data

//TODO: maybe there is a better way to store these timestamps
type Rental struct {
	RentalID     int
	RentedFrom   string
	RentedTill   string
	CreatedOn    string
	LastModified string
	Renter       BasicRenter
	Trackables   []BasicTrackable
}
