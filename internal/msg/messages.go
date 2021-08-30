package msg

import (
	"errors"
	"fmt"
	"time"
)

const (
	// Error ...
	Error = "\x1b[1mError:\x1b[0m %v"
	// Response ...
	Response = "\x1b[1mResponse:\x1b[0m %v"
)

const (
	// StrTemplateTicket ...
	StrTemplateTicket = "%v | %s"
	// TimeLayout ...
	TimeLayout = "2006-01-02T15:04:05.9999999Z"
)

// Fail ..
func Fail(s string) (err error) {
	err = errors.New(fmt.Sprintf(Error, s))
	return
}

// TimeNowFormatted ...
func TimeNowFormatted() (rtn string) {
	t := time.Now()
	// apply formatting
	rtn = t.UTC().Format(TimeLayout)
	return
}
