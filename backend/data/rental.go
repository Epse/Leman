// This package provides a bunch of data structs and necessary functions
package data

import (
	"time"
)

//TODO: maybe there is a better way to store these timestamps
type Rental struct {
	RentalID     int
	RentedFrom   time.Time
	RentedTill   time.Time
	CreatedOn    time.Time
	LastModified time.Time
	Renter       BasicRenter
	Trackables   []BasicTrackable
}
