package types

const (
	// BYN ...
	BYN Currency = 54
	// USD ...
	USD Currency = 251
	// EUR ...
	EUR Currency = 93
	// RUB ...
	RUB Currency = 209
)

// Currency ...
type Currency int

// IsValid check for currency input
// 	Can check if currency is allowed in a system
func (c Currency) IsValid() bool {

	switch c {
	case BYN, USD:
		return true
	}
	return false
}
