package types

import "unicode/utf8"

// Description ...
type Description string

func (d Description) String() string {
	return string(d)
}

// IsValid ..
// 	len <= 1000 letters
func (d Description) IsValid() bool {
	lengthChars := utf8.RuneCountInString
	var lengthMax = 1000
	if lengthChars(d.String()) > lengthMax {
		return false
	}
	return true
}
