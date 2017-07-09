package clistruct

import (
	"fmt"
	"reflect"
)

// ErrInvalid is an error indicating that invalid values was passed.
type ErrInvalid struct {
	v interface{}
}

func (e *ErrInvalid) Error() string {
	return fmt.Sprintf(
		"Reflect reports this value is invalid '%#v'",
		e.v,
	)
}

// NewErrInvalid creates new ErrInvalid.
func NewErrInvalid(v interface{}) error {
	return &ErrInvalid{v}
}

//

// ErrInvalidKind is an error indicating that reflect.Kind of
// value is not expected in the context it was used.
type ErrInvalidKind struct {
	expected reflect.Kind
	got      reflect.Kind
}

func (e *ErrInvalidKind) Error() string {
	return fmt.Sprintf(
		"Expected '%s' kind, got '%s'",
		e.expected,
		e.got,
	)
}

// NewErrInvalidKind creates new ErrInvalidKind.
func NewErrInvalidKind(expected, got reflect.Kind) error {
	return &ErrInvalidKind{expected, got}
}

//

// ErrPtrRequired is an error indicating that a pointer entity required.
type ErrPtrRequired struct {
	v interface{}
}

func (e *ErrPtrRequired) Error() string {
	return fmt.Sprintf(
		"A pointer to the value '%#v' is required, not the value itself",
		e.v,
	)
}

// NewErrPtrRequired creates new ErrPtrRequired.
func NewErrPtrRequired(v interface{}) error {
	return &ErrPtrRequired{v}
}

//

// ErrTypeMistmatch is an error indicating that a wanted type
// is not equal to actual.
type ErrTypeMistmatch struct {
	want string
	got  string
}

func (e *ErrTypeMistmatch) Error() string {
	return fmt.Sprintf(
		"Type mistmatch, want '%s', got '%s'",
		e.want,
		e.got,
	)
}

// NewErrTypeMistmatch creates new ErrTypeMistmatch.
func NewErrTypeMistmatch(want string, got string) error {
	return &ErrTypeMistmatch{want, got}
}

//

// ErrFlagTypeCanNotHaveValue is an error indicating that
// flag type with specified name takes no value.
type ErrFlagTypeCanNotHaveValue struct {
	t string
}

func (e *ErrFlagTypeCanNotHaveValue) Error() string {
	return fmt.Sprintf(
		"Flag of type '%s' can not have value",
		e.t,
	)
}

// NewErrFlagTypeCanNotHaveValue creates new ErrFlagTypeCanNotHaveValue.
func NewErrFlagTypeCanNotHaveValue(t string) error {
	return &ErrFlagTypeCanNotHaveValue{t}
}
