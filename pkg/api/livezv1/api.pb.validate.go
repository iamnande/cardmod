// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: iamnande/cardmod/livez/v1/api.proto

package livezv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DescribeLivezRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DescribeLivezRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DescribeLivezRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DescribeLivezRequestMultiError, or nil if none found.
func (m *DescribeLivezRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DescribeLivezRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DescribeLivezRequestMultiError(errors)
	}

	return nil
}

// DescribeLivezRequestMultiError is an error wrapping multiple validation
// errors returned by DescribeLivezRequest.ValidateAll() if the designated
// constraints aren't met.
type DescribeLivezRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DescribeLivezRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DescribeLivezRequestMultiError) AllErrors() []error { return m }

// DescribeLivezRequestValidationError is the validation error returned by
// DescribeLivezRequest.Validate if the designated constraints aren't met.
type DescribeLivezRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeLivezRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeLivezRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeLivezRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeLivezRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeLivezRequestValidationError) ErrorName() string {
	return "DescribeLivezRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeLivezRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeLivezRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeLivezRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeLivezRequestValidationError{}

// Validate checks the field values on DescribeLivezResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DescribeLivezResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DescribeLivezResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DescribeLivezResponseMultiError, or nil if none found.
func (m *DescribeLivezResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DescribeLivezResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return DescribeLivezResponseMultiError(errors)
	}

	return nil
}

// DescribeLivezResponseMultiError is an error wrapping multiple validation
// errors returned by DescribeLivezResponse.ValidateAll() if the designated
// constraints aren't met.
type DescribeLivezResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DescribeLivezResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DescribeLivezResponseMultiError) AllErrors() []error { return m }

// DescribeLivezResponseValidationError is the validation error returned by
// DescribeLivezResponse.Validate if the designated constraints aren't met.
type DescribeLivezResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeLivezResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeLivezResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeLivezResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeLivezResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeLivezResponseValidationError) ErrorName() string {
	return "DescribeLivezResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeLivezResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeLivezResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeLivezResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeLivezResponseValidationError{}
