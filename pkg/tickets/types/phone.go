package types

import (
	"regexp"
)

// Phone ...
type Phone string

func (p Phone) String() string {
	return string(p)
}

// IsValid formatting examples:
// 	123-456-7890, (123) 456-7890, 123 456 7890
// 	123.456.7890, +91 (123) 456-7890
func (p Phone) IsValid() bool {
	re := regexp.MustCompile(`(^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$)`)
	subMatch := re.FindStringSubmatch(p.String())
	if len(subMatch) < 2 {
		return false
	}
	return true
}
