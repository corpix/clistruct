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
	"reflect"
	"time"

	"github.com/urfave/cli"
)

const (
	fieldTag = "cli"
	nameTag  = "name"
	typeTag  = "type"
	usageTag = "usage"
	valueTag = "value"
)

const (
	boolType        = "bool"
	boolTType       = "boolt"
	float64Type     = "float64"
	intType         = "int"
	int64Type       = "int64"
	intSliceType    = "intslice"
	int64SliceType  = "int64slice"
	stringType      = "string"
	stringSliceType = "stringslice"
	uintType        = "uint"
	uint64Type      = "uint64"
	durationType    = "duration"
	genericType     = "generic"
)

var (
	fieldToFlag = map[string]cli.Flag{
		typeName(new(bool)):          new(cli.BoolFlag),
		typeName(new(float64)):       new(cli.Float64Flag),
		typeName(new(int)):           new(cli.IntFlag),
		typeName(new(int64)):         new(cli.Int64Flag),
		typeName(new([]int)):         new(cli.IntSliceFlag),
		typeName(new([]int64)):       new(cli.Int64SliceFlag),
		typeName(new(string)):        new(cli.StringFlag),
		typeName(new([]string)):      new(cli.StringSliceFlag),
		typeName(new(uint)):          new(cli.UintFlag),
		typeName(new(uint64)):        new(cli.Uint64Flag),
		typeName(new(time.Duration)): new(cli.DurationFlag),
	}

	fieldTypeTagToFlag = map[string]cli.Flag{
		boolType:        new(cli.BoolFlag),
		boolTType:       new(cli.BoolTFlag),
		float64Type:     new(cli.Float64Flag),
		intType:         new(cli.IntFlag),
		int64Type:       new(cli.Int64Flag),
		intSliceType:    new(cli.IntSliceFlag),
		int64SliceType:  new(cli.Int64SliceFlag),
		stringType:      new(cli.StringFlag),
		stringSliceType: new(cli.StringSliceFlag),
		uintType:        new(cli.UintFlag),
		uint64Type:      new(cli.Uint64Flag),
		durationType:    new(cli.DurationFlag),
		genericType:     new(cli.GenericFlag),
	}
)

// FlagsFromStruct generates cli.Flag slice for github.com/urfave/cli
// from the struct fields.
func FlagsFromStruct(v interface{}) ([]cli.Flag, error) {
	err := checkValue(v)
	if err != nil {
		return nil, err
	}

	return flagsFromStruct(v)
}

func flagsFromStruct(v interface{}) ([]cli.Flag, error) {
	var (
		reflectType  = indirectType(reflect.TypeOf(v))
		reflectValue = indirectValue(reflect.ValueOf(v))
		flag         cli.Flag
		flags        []cli.Flag
		field        reflect.StructField
		err          error
	)

	if !reflectValue.IsValid() {
		return nil, NewErrInvalid(v)
	}

	err = shouldBeStruct(reflectValue)
	if err != nil {
		return nil, err
	}

	for n := 0; n < reflectValue.NumField(); n++ {
		field = reflectType.Field(n)
		if !isStructFieldExported(field) {
			continue
		}

		flag, err = flagFromStructField(field)
		if err != nil {
			return nil, err
		}

		flags = append(
			flags,
			flag,
		)
	}

	return flags, nil
}

func newFlagFromStructField(field reflect.StructField) cli.Flag {
	var (
		t cli.Flag
	)

	t = fieldTypeTagToFlag[field.Tag.Get(typeTag)]
	if t == nil {
		t = fieldToFlag[field.Type.String()]
	}
	if t == nil {
		t = fieldTypeTagToFlag[genericType]
	}

	return reflect.New(reflect.TypeOf(t).Elem()).
		Interface().(cli.Flag)
}

func flagFromStructField(field reflect.StructField) (cli.Flag, error) {
	var (
		flag cli.Flag
		err  error
	)

	flag = newFlagFromStructField(field)

	err = setStructField(flag, "Name", field.Tag.Get(nameTag))
	if err != nil {
		return nil, err
	}
	err = setStructField(flag, "Usage", field.Tag.Get(usageTag))
	if err != nil {
		return nil, err
	}
	// FIXME: Values should be parsed
	// err = setStructField(flag, "Value", field.Tag.Get(valueTag))
	// if err != nil {
	// 	return nil, err
	// }

	return flag, nil
}

// FlagsToStruct folds a flags from context into the struct fields in v.
func FlagsToStruct(context *cli.Context, v interface{}) error {
	err := checkValue(v)
	if err != nil {
		return err
	}

	return nil
}

//

func setStructField(v interface{}, fieldName string, value interface{}) error {
	field := indirectValue(reflect.ValueOf(v)).
		FieldByName(fieldName)

	if !field.IsValid() || !field.CanSet() {
		return NewErrInvalid(v)
	}

	field.Set(reflect.ValueOf(value))
	return nil
}

func checkValue(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return NewErrPtrRequired(v)
	}

	return nil
}

func shouldBeStruct(reflectValue reflect.Value) error {
	if reflectValue.Kind() != reflect.Struct {
		return NewErrInvalidKind(
			reflect.Struct,
			reflectValue.Kind(),
		)
	}

	return nil
}

func isStructFieldExported(field reflect.StructField) bool {
	// From reflect docs:
	// PkgPath is the package path that qualifies a lower case (unexported)
	// field name. It is empty for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	return field.PkgPath == ""
}

func indirectValue(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		return reflectValue.Elem()
	}
	return reflectValue
}

func indirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		return reflectType.Elem()
	}
	return reflectType
}

func typeName(v interface{}) string {
	return indirectType(reflect.TypeOf(v)).String()
}
