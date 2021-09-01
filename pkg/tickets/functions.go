package tickets

import (
	"unicode/utf8"

	"github.com/go-playground/validator"

	"github.com/rodkevich/tbd/internal/msg"
)

var (
	lengthNameMax  = 200
	lengthLinksMax = 3
)

var (
	valErrorCommon       = msg.Fail(`validation failed`)
	valErrorName         = msg.Fail(`validation failed: Name`)
	valErrorDescription  = msg.Fail(`validation failed: Description`)
	valErrorLinkOverflow = msg.Fail(`validation failed: PhotoLinks overflow. Limit = 3`)
	valErrorHardcoded    = msg.Fail(`validation that cannot be turned-off failed: PhoneNumber`)
)

// validationConfig ...
type validationConfig struct {
	StructTags  bool
	Name        bool
	PhotoLinks  bool
	Description bool
}

// ValidationOption ...
type ValidationOption func(*validationConfig)

// WithoutStructTags do NOT use validation with JSON struct tags
func WithoutStructTags() ValidationOption {
	return func(v *validationConfig) {
		v.StructTags = false
	}
}

// WithNameCheck ...
func WithNameCheck() ValidationOption {
	return func(v *validationConfig) {
		v.Name = true
	}
}

// WithPhotoLinksCheck ...
func WithPhotoLinksCheck() ValidationOption {
	return func(v *validationConfig) {
		v.PhotoLinks = true
	}
}

// WithDescriptionCheck ...
func WithDescriptionCheck() ValidationOption {
	return func(v *validationConfig) {
		v.Description = true
	}
}

// TicketValidation ...
func TicketValidation(t Ticket, opts ...ValidationOption) (err error) {

	const (
		// enabled by default:
		defaultStructTags = true
		// optional:
		defaultName        = false
		defaultPhotoLinks  = false
		defaultDescription = false
	)

	enabled := &validationConfig{
		StructTags:  defaultStructTags,
		Name:        defaultName,
		PhotoLinks:  defaultPhotoLinks,
		Description: defaultDescription,
	}

	for _, opt := range opts {
		opt(enabled)
	}

	// first way - validate through struct tags:
	if enabled.StructTags {
		validate := validator.New()
		err = validate.Struct(t)
		if err != nil {
			_ = err.(validator.ValidationErrors)
			return valErrorCommon
		}
	}

	// another way - with self methods:
	if enabled.Name {
		length := utf8.RuneCountInString
		if length(t.Name) > lengthNameMax {
			return valErrorName
		}
	}

	if enabled.PhotoLinks {
		if len(t.PhotoLinks) > lengthLinksMax {
			return valErrorLinkOverflow
		}

		// for x, link := range t.PhotoLinks {
		// 	if !link.IsValid() {
		// 		message := `validation failed: validationErrorPhotoLinks N:` + strconv.Itoa(x+1)
		// 		return msg.Fail(message)
		// 	}
		// }
	}

	if enabled.Description && !t.Description.IsValid() {
		return valErrorDescription
	}

	// cannot be turned-off:
	if !t.PhoneNumber.IsValid() {
		return valErrorHardcoded
	}
	return
}
