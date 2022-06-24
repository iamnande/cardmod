// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: iamnande/cardmod/calculate/v1/api.proto

package calculatev1

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

// Validate checks the field values on CalculateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CalculateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CalculateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CalculateRequestMultiError, or nil if none found.
func (m *CalculateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CalculateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTarget()); l < 3 || l > 25 {
		err := CalculateRequestValidationError{
			field:  "Target",
			reason: "value length must be between 3 and 25 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CalculateRequest_Target_Pattern.MatchString(m.GetTarget()) {
		err := CalculateRequestValidationError{
			field:  "Target",
			reason: "value does not match regex pattern \"^[-, \\\\w]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetCount(); val < 1 || val > 300 {
		err := CalculateRequestValidationError{
			field:  "Count",
			reason: "value must be inside range [1, 300]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CalculateRequestMultiError(errors)
	}

	return nil
}

// CalculateRequestMultiError is an error wrapping multiple validation errors
// returned by CalculateRequest.ValidateAll() if the designated constraints
// aren't met.
type CalculateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalculateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalculateRequestMultiError) AllErrors() []error { return m }

// CalculateRequestValidationError is the validation error returned by
// CalculateRequest.Validate if the designated constraints aren't met.
type CalculateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalculateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalculateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalculateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalculateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalculateRequestValidationError) ErrorName() string { return "CalculateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CalculateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalculateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalculateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalculateRequestValidationError{}

var _CalculateRequest_Target_Pattern = regexp.MustCompile("^[-, \\w]+$")

// Validate checks the field values on CalculateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CalculateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CalculateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CalculateResponseMultiError, or nil if none found.
func (m *CalculateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CalculateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRefinements() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalculateResponseValidationError{
						field:  fmt.Sprintf("Refinements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalculateResponseValidationError{
						field:  fmt.Sprintf("Refinements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalculateResponseValidationError{
					field:  fmt.Sprintf("Refinements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CalculateResponseMultiError(errors)
	}

	return nil
}

// CalculateResponseMultiError is an error wrapping multiple validation errors
// returned by CalculateResponse.ValidateAll() if the designated constraints
// aren't met.
type CalculateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalculateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalculateResponseMultiError) AllErrors() []error { return m }

// CalculateResponseValidationError is the validation error returned by
// CalculateResponse.Validate if the designated constraints aren't met.
type CalculateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalculateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalculateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalculateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalculateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalculateResponseValidationError) ErrorName() string {
	return "CalculateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CalculateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalculateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalculateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalculateResponseValidationError{}

// Validate checks the field values on CalculateResponse_Refinement with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CalculateResponse_Refinement) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CalculateResponse_Refinement with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CalculateResponse_RefinementMultiError, or nil if none found.
func (m *CalculateResponse_Refinement) ValidateAll() error {
	return m.validate(true)
}

func (m *CalculateResponse_Refinement) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Source

	// no validation rules for Count

	for idx, item := range m.GetRefinements() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CalculateResponse_RefinementValidationError{
						field:  fmt.Sprintf("Refinements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CalculateResponse_RefinementValidationError{
						field:  fmt.Sprintf("Refinements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalculateResponse_RefinementValidationError{
					field:  fmt.Sprintf("Refinements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CalculateResponse_RefinementMultiError(errors)
	}

	return nil
}

// CalculateResponse_RefinementMultiError is an error wrapping multiple
// validation errors returned by CalculateResponse_Refinement.ValidateAll() if
// the designated constraints aren't met.
type CalculateResponse_RefinementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CalculateResponse_RefinementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CalculateResponse_RefinementMultiError) AllErrors() []error { return m }

// CalculateResponse_RefinementValidationError is the validation error returned
// by CalculateResponse_Refinement.Validate if the designated constraints
// aren't met.
type CalculateResponse_RefinementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalculateResponse_RefinementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalculateResponse_RefinementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalculateResponse_RefinementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalculateResponse_RefinementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalculateResponse_RefinementValidationError) ErrorName() string {
	return "CalculateResponse_RefinementValidationError"
}

// Error satisfies the builtin error interface
func (e CalculateResponse_RefinementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalculateResponse_Refinement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalculateResponse_RefinementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalculateResponse_RefinementValidationError{}