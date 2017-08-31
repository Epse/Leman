// This package provides a bunch of data structs and necessary functions
package data

import (
	"time"
)

type User struct {
	UserID       int
	FirstName    string
	FamilyName   string
	PasswordHash string
	Role         BasicRole
	Email        string
	Phone        string
	CreatedOn    time.Time
	LastLogin    time.Time
}
