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
	"strconv"
	"strings"
)

const (
	fieldTag = "cli"
	nameTag  = "name"
	typeTag  = "type"
	usageTag = "usage"
	valueTag = "value"
)

const (
	listDelimiter = ","
)

const (
	boolTypeTag        = "bool"
	boolTTypeTag       = "boolt"
	float64TypeTag     = "float64"
	intTypeTag         = "int"
	int64TypeTag       = "int64"
	intSliceTypeTag    = "intslice"
	int64SliceTypeTag  = "int64slice"
	stringTypeTag      = "string"
	stringSliceTypeTag = "stringslice"
	uintTypeTag        = "uint"
	uint64TypeTag      = "uint64"
	durationTypeTag    = "duration"
	genericTypeTag     = "generic"
)

var (
	boolType        = typeName(*new(bool))
	uintType        = typeName(*new(uint))
	uint64Type      = typeName(*new(uint64))
	intType         = typeName(*new(int))
	int64Type       = typeName(*new(int64))
	float64Type     = typeName(*new(float64))
	intSliceType    = typeName(*new([]int))
	int64SliceType  = typeName(*new([]int64))
	stringType      = typeName(*new(string))
	stringSliceType = typeName(*new([]string))
	durationType    = typeName(*new(time.Duration))
)

var (
	uintValueType        = uintType
	uint64ValueType      = uint64Type
	intValueType         = intType
	int64ValueType       = int64Type
	float64ValueType     = float64Type
	intSliceValueType    = typeName(new(cli.IntSlice))
	int64SliceValueType  = typeName(new(cli.Int64Slice))
	stringValueType      = stringType
	stringSliceValueType = typeName(new(cli.StringSlice))
	durationValueType    = durationType
)

var (
	typeToFlag = map[string]cli.Flag{
		boolType:        new(cli.BoolFlag),
		uintType:        new(cli.UintFlag),
		uint64Type:      new(cli.Uint64Flag),
		intType:         new(cli.IntFlag),
		int64Type:       new(cli.Int64Flag),
		float64Type:     new(cli.Float64Flag),
		intSliceType:    new(cli.IntSliceFlag),
		int64SliceType:  new(cli.Int64SliceFlag),
		stringType:      new(cli.StringFlag),
		stringSliceType: new(cli.StringSliceFlag),
		durationType:    new(cli.DurationFlag),
	}

	typeTagToFlag = map[string]cli.Flag{
		boolTypeTag:        *new(cli.BoolFlag),
		boolTTypeTag:       *new(cli.BoolTFlag),
		float64TypeTag:     *new(cli.Float64Flag),
		intTypeTag:         *new(cli.IntFlag),
		int64TypeTag:       *new(cli.Int64Flag),
		intSliceTypeTag:    *new(cli.IntSliceFlag),
		int64SliceTypeTag:  *new(cli.Int64SliceFlag),
		stringTypeTag:      *new(cli.StringFlag),
		stringSliceTypeTag: *new(cli.StringSliceFlag),
		uintTypeTag:        *new(cli.UintFlag),
		uint64TypeTag:      *new(cli.Uint64Flag),
		durationTypeTag:    *new(cli.DurationFlag),
		genericTypeTag:     *new(cli.GenericFlag),
	}

	flagsWithoutValues = map[string]bool{
		boolType: true,
	}

	valueFromString = map[string]func(string) (interface{}, error){
		uintValueType: func(v string) (interface{}, error) {
			u, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			return uint(u), nil
		},
		uint64ValueType: func(v string) (interface{}, error) {
			return strconv.ParseUint(v, 10, 64)
		},
		intValueType: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return nil, err
			}
			return int(i), nil
		},
		int64ValueType: func(v string) (interface{}, error) {
			return strconv.ParseInt(v, 10, 64)
		},
		float64ValueType: func(v string) (interface{}, error) {
			return strconv.ParseFloat(v, 64)
		},
		intSliceValueType: func(v string) (interface{}, error) {
			var (
				ints     = strings.Split(v, listDelimiter)
				intSlice = make(cli.IntSlice, len(ints))
				i        int64
				err      error
			)

			for k, v := range ints {
				i, err = strconv.ParseInt(v, 10, 32)
				if err != nil {
					return nil, err
				}
				intSlice[k] = int(i)
			}

			return &intSlice, nil
		},
		int64SliceValueType: func(v string) (interface{}, error) {
			var (
				ints       = strings.Split(v, listDelimiter)
				int64Slice = make(cli.Int64Slice, len(ints))
				i          int64
				err        error
			)

			for k, v := range ints {
				i, err = strconv.ParseInt(v, 10, 64)
				if err != nil {
					return nil, err
				}
				int64Slice[k] = i
			}

			return &int64Slice, nil
		},
		stringValueType: func(v string) (interface{}, error) {
			return v, nil
		},
		stringSliceValueType: func(v string) (interface{}, error) {
			stringSlice := cli.StringSlice(strings.Split(v, listDelimiter))
			return &stringSlice, nil
		},
		durationValueType: func(v string) (interface{}, error) {
			return time.ParseDuration(v)
		},
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
			indirectValue(reflect.ValueOf(flag)).
				Interface().(cli.Flag),
		)
	}

	return flags, nil
}

func parseValueFromString(v string, targetType reflect.Type) (interface{}, error) {
	parser, ok := valueFromString[targetType.String()]
	if ok {
		return parser(v)
	}

	return v, nil
}

func newFlagFromStructField(field reflect.StructField) cli.Flag {
	var (
		t cli.Flag
	)

	t = typeTagToFlag[field.Tag.Get(typeTag)]
	if t == nil {
		t = typeToFlag[field.Type.String()]
	}
	if t == nil {
		t = typeTagToFlag[genericTypeTag]
	}

	return reflect.
		New(
			indirectType(
				reflect.TypeOf(t),
			),
		).
		Interface().(cli.Flag)
}

func flagFromStructField(field reflect.StructField) (cli.Flag, error) {
	var (
		flag       cli.Flag
		valueField reflect.Value
		value      interface{}
		err        error
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

	valueString := field.Tag.Get(valueTag)
	if valueString != "" && !flagsWithoutValues[field.Type.String()] {
		valueField, err = getStructField(flag, "Value")
		if err != nil {
			return nil, err
		}
		value, err = parseValueFromString(
			valueString,
			valueField.Type(),
		)
		if err != nil {
			return nil, err
		}

		err = setStructField(flag, "Value", value)
		if err != nil {
			return nil, err
		}
	}

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
