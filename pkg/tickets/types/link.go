package types

import "net/url"

// Link ...
type Link string

func (l Link) String() string {
	return string(l)
}

// IsValid ...
func (l Link) IsValid() bool {
	u, err := url.Parse(l.String())
	return err == nil && u.Scheme != "" && u.Host != ""
}
