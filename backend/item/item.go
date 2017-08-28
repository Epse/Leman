// Item provides a basic Item struct and some basic functions for operating on them.
package item

// Item is the basic struct for defining an item in inventory, whether individually or group tracked.
// A property indentified with the `IGT` suffix is only used if this item is group tracked.
// The suffix `IIT` is used for individually tracked items.
type Item struct {
	ID                    uint64
	Name                  string
	IsIndividuallyTracked bool
	TotalQuantityIGT      uint64
	QuantityAvailableIGT  uint64
	CurrentLocationIDIIT  uint64
	Brand                 string
	PricePerTime          float32
	TimeUnit              string
	inStockIIT            bool
}

func (i Item) IsInStock() bool {
	if i.IsIndividuallyTracked {
		return i.InStockIIT
	} else {
		return i.QuantityAvailableIGT > 0
	}
}
