package clistruct

// The MIT License (MIT)
//
// Copyright © 2017 Dmitry Moskowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
