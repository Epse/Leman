// This package provides some boilerplate for the API
package api

import (
	"encoding/json"
	"github.com/Epse/Leman/backend/data"
)

// The ListResponse makes up the basic framework for a response that contains a list of items.
type ListResponse struct {
	Products []data.BasicProduct
}

// The GenerateJSON function basically just calls the json.Marshal() function, but is clearer to use
func (lr ListResponse) GenerateJSON() ([]byte, error) {
	return json.Marshal(lr)
}

// The ItemResponse struct currently has no additional value over just using json.Marshal() on an item.Item directly, but allows for future additions and is clearer.
type ProductResponse struct {
	Product data.BasicProduct
}

// The GenerateJSON function basically just calls the json.Marshal() function, but is clearer to use
func (ir ProductResponse) GenerateJSON() ([]byte, error) {
	return json.Marshal(ir)
}
