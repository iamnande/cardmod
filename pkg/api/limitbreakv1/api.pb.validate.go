// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: iamnande/cardmod/limitbreak/v1/api.proto

package limitbreakv1

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

// Validate checks the field values on GetLimitBreakRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLimitBreakRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLimitBreakRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLimitBreakRequestMultiError, or nil if none found.
func (m *GetLimitBreakRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLimitBreakRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 25 {
		err := GetLimitBreakRequestValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 25 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_GetLimitBreakRequest_Name_Pattern.MatchString(m.GetName()) {
		err := GetLimitBreakRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[-, \\\\w]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetLimitBreakRequestMultiError(errors)
	}
	return nil
}

// GetLimitBreakRequestMultiError is an error wrapping multiple validation
// errors returned by GetLimitBreakRequest.ValidateAll() if the designated
// constraints aren't met.
type GetLimitBreakRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLimitBreakRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLimitBreakRequestMultiError) AllErrors() []error { return m }

// GetLimitBreakRequestValidationError is the validation error returned by
// GetLimitBreakRequest.Validate if the designated constraints aren't met.
type GetLimitBreakRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLimitBreakRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLimitBreakRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLimitBreakRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLimitBreakRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLimitBreakRequestValidationError) ErrorName() string {
	return "GetLimitBreakRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetLimitBreakRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLimitBreakRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLimitBreakRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLimitBreakRequestValidationError{}

var _GetLimitBreakRequest_Name_Pattern = regexp.MustCompile("^[-, \\w]+$")

// Validate checks the field values on ListLimitBreaksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListLimitBreaksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListLimitBreaksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListLimitBreaksRequestMultiError, or nil if none found.
func (m *ListLimitBreaksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListLimitBreaksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListLimitBreaksRequestMultiError(errors)
	}
	return nil
}

// ListLimitBreaksRequestMultiError is an error wrapping multiple validation
// errors returned by ListLimitBreaksRequest.ValidateAll() if the designated
// constraints aren't met.
type ListLimitBreaksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListLimitBreaksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListLimitBreaksRequestMultiError) AllErrors() []error { return m }

// ListLimitBreaksRequestValidationError is the validation error returned by
// ListLimitBreaksRequest.Validate if the designated constraints aren't met.
type ListLimitBreaksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListLimitBreaksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListLimitBreaksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListLimitBreaksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListLimitBreaksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListLimitBreaksRequestValidationError) ErrorName() string {
	return "ListLimitBreaksRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListLimitBreaksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListLimitBreaksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListLimitBreaksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListLimitBreaksRequestValidationError{}

// Validate checks the field values on ListLimitBreaksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListLimitBreaksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListLimitBreaksResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListLimitBreaksResponseMultiError, or nil if none found.
func (m *ListLimitBreaksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListLimitBreaksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLimitbreaks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListLimitBreaksResponseValidationError{
						field:  fmt.Sprintf("Limitbreaks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListLimitBreaksResponseValidationError{
						field:  fmt.Sprintf("Limitbreaks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListLimitBreaksResponseValidationError{
					field:  fmt.Sprintf("Limitbreaks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListLimitBreaksResponseMultiError(errors)
	}
	return nil
}

// ListLimitBreaksResponseMultiError is an error wrapping multiple validation
// errors returned by ListLimitBreaksResponse.ValidateAll() if the designated
// constraints aren't met.
type ListLimitBreaksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListLimitBreaksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListLimitBreaksResponseMultiError) AllErrors() []error { return m }

// ListLimitBreaksResponseValidationError is the validation error returned by
// ListLimitBreaksResponse.Validate if the designated constraints aren't met.
type ListLimitBreaksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListLimitBreaksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListLimitBreaksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListLimitBreaksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListLimitBreaksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListLimitBreaksResponseValidationError) ErrorName() string {
	return "ListLimitBreaksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListLimitBreaksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListLimitBreaksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListLimitBreaksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListLimitBreaksResponseValidationError{}