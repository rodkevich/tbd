package types

const (
	// BYN ...
	BYN Currency = 54
	// USD ...
	USD Currency = 251
	// EUR ...
	EUR Currency = 93
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

// CurrencyToString ...
func CurrencyToString(c Currency) string {
	switch c {
	case BYN:
		return "Candy wrappers"
	case USD:
		return "Money"
	case EUR:
		return "Not yet supported"
	}
	return "ErrorCase"
}
