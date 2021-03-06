package clistruct

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestFlagsFromStructWithTags(t *testing.T) {
	type custom struct{}

	sample := struct {
		Bool        bool          `name:"bool"        type:"bool"        usage:"hello"`
		BoolT       bool          `name:"boolt"       type:"boolt"       usage:"hello"`
		UInt        uint          `name:"uint"        type:"uint"        usage:"hello" value:"1"`
		UInt64      uint64        `name:"uint64"      type:"uint64"      usage:"hello" value:"1"`
		Int         int           `name:"int"         type:"int"         usage:"hello" value:"1"`
		Int64       int64         `name:"int64"       type:"int64"       usage:"hello" value:"-1"`
		Float64     float64       `name:"float64"     type:"float64"     usage:"hello" value:"1.5"`
		IntSlice    []int         `name:"intslice"    type:"intslice"    usage:"hello" value:"1,2,3,-1"`
		Int64Slice  []int64       `name:"int64slice"  type:"int64slice"  usage:"hello" value:"1,2,3,-1"`
		String      string        `name:"string"      type:"string"      usage:"hello" value:"some string"`
		StringSlice []string      `name:"stringslice" type:"stringslice" usage:"hello" value:"some,string,slice"`
		Duration    time.Duration `name:"duration"    type:"duration"    usage:"hello" value:"2h1m10s"`
		Custom      custom        `name:"custom"      type:"generic"     usage:"hello"`
	}{}
	flags := []cli.Flag{
		cli.BoolFlag{Name: "bool", Usage: "hello"},
		cli.BoolTFlag{Name: "boolt", Usage: "hello"},
		cli.UintFlag{Name: "uint", Usage: "hello", Value: uint(1)},
		cli.Uint64Flag{Name: "uint64", Usage: "hello", Value: uint64(1)},
		cli.IntFlag{Name: "int", Usage: "hello", Value: 1},
		cli.Int64Flag{Name: "int64", Usage: "hello", Value: int64(-1)},
		cli.Float64Flag{Name: "float64", Usage: "hello", Value: float64(1.5)},
		cli.IntSliceFlag{Name: "intslice", Usage: "hello", Value: &cli.IntSlice{1, 2, 3, -1}},
		cli.Int64SliceFlag{Name: "int64slice", Usage: "hello", Value: &cli.Int64Slice{1, 2, 3, -1}},
		cli.StringFlag{Name: "string", Usage: "hello", Value: "some string"},
		cli.StringSliceFlag{Name: "stringslice", Usage: "hello", Value: &cli.StringSlice{"some", "string", "slice"}},
		cli.DurationFlag{Name: "duration", Usage: "hello", Value: (2 * 60 * time.Minute) + (time.Minute) + (10 * time.Second)},
		cli.GenericFlag{Name: "custom", Usage: "hello"},
	}

	result, err := FlagsFromStruct(&sample)
	if err != nil {
		t.Error(err)
		return
	}
	assert.EqualValues(t, flags, result)
}

func TestFlagsFromStructWithoutNameTags(t *testing.T) {
	type custom struct{}

	sample := struct {
		Bool        bool          `type:"bool"        usage:"hello"`
		BoolT       bool          `type:"boolt"       usage:"hello"`
		UInt        uint          `type:"uint"        usage:"hello" value:"1"`
		UInt64      uint64        `type:"uint64"      usage:"hello" value:"1"`
		Int         int           `type:"int"         usage:"hello" value:"1"`
		Int64       int64         `type:"int64"       usage:"hello" value:"-1"`
		Float64     float64       `type:"float64"     usage:"hello" value:"1.5"`
		IntSlice    []int         `type:"intslice"    usage:"hello" value:"1,2,3,-1"`
		Int64Slice  []int64       `type:"int64slice"  usage:"hello" value:"1,2,3,-1"`
		String      string        `type:"string"      usage:"hello" value:"some string"`
		StringSlice []string      `type:"stringslice" usage:"hello" value:"some,string,slice"`
		Duration    time.Duration `type:"duration"    usage:"hello" value:"2h1m10s"`
		Custom      custom        `type:"generic"     usage:"hello"`
	}{}
	flags := []cli.Flag{
		cli.BoolFlag{Name: "bool", Usage: "hello"},
		cli.BoolTFlag{Name: "boolt", Usage: "hello"},
		cli.UintFlag{Name: "uint", Usage: "hello", Value: uint(1)},
		cli.Uint64Flag{Name: "uint64", Usage: "hello", Value: uint64(1)},
		cli.IntFlag{Name: "int", Usage: "hello", Value: 1},
		cli.Int64Flag{Name: "int64", Usage: "hello", Value: int64(-1)},
		cli.Float64Flag{Name: "float64", Usage: "hello", Value: float64(1.5)},
		cli.IntSliceFlag{Name: "intslice", Usage: "hello", Value: &cli.IntSlice{1, 2, 3, -1}},
		cli.Int64SliceFlag{Name: "int64slice", Usage: "hello", Value: &cli.Int64Slice{1, 2, 3, -1}},
		cli.StringFlag{Name: "string", Usage: "hello", Value: "some string"},
		cli.StringSliceFlag{Name: "stringslice", Usage: "hello", Value: &cli.StringSlice{"some", "string", "slice"}},
		cli.DurationFlag{Name: "duration", Usage: "hello", Value: (2 * 60 * time.Minute) + (time.Minute) + (10 * time.Second)},
		cli.GenericFlag{Name: "custom", Usage: "hello"},
	}

	result, err := FlagsFromStruct(&sample)
	if err != nil {
		t.Error(err)
		return
	}
	assert.EqualValues(t, flags, result)
}

func TestFlagsFromStructBoolHasNoValue(t *testing.T) {
	sample := struct {
		Bool bool `name:"bool"        type:"bool"        usage:"hello" value:"true"`
	}{}

	result, err := FlagsFromStruct(&sample)
	assert.NotNil(t, err)
	switch err.(type) {
	case *ErrFlagTypeCanNotHaveValue:
	default:
		t.Error(err)
		return
	}
	assert.EqualValues(t, ([]cli.Flag)(nil), result)
}

func TestFlagsFromStructBoolThasNoValue(t *testing.T) {
	sample := struct {
		BoolT bool `name:"boolt"       type:"boolt"       usage:"hello" value:"true"`
	}{}

	result, err := FlagsFromStruct(&sample)
	assert.NotNil(t, err)
	switch err.(type) {
	case *ErrFlagTypeCanNotHaveValue:
	default:
		t.Error(err)
		return
	}
	assert.EqualValues(t, ([]cli.Flag)(nil), result)
}

func TestFlagsRoStructWithTags(t *testing.T) {
	// FIXME: generic types is nil in .cli for some reason oO
	//type custom struct{}
	type Sample struct {
		Bool        bool          `name:"bool"        type:"bool"        usage:"hello"`
		BoolT       bool          `name:"boolt"       type:"boolt"       usage:"hello"`
		UInt        uint          `name:"uint"        type:"uint"        usage:"hello" value:"1"`
		UInt64      uint64        `name:"uint64"      type:"uint64"      usage:"hello" value:"1"`
		Int         int           `name:"int"         type:"int"         usage:"hello" value:"1"`
		Int64       int64         `name:"int64"       type:"int64"       usage:"hello" value:"-1"`
		Float64     float64       `name:"float64"     type:"float64"     usage:"hello" value:"1.5"`
		IntSlice    []int         `name:"intslice"    type:"intslice"    usage:"hello" value:"1,2,3,-1"`
		Int64Slice  []int64       `name:"int64slice"  type:"int64slice"  usage:"hello" value:"1,2,3,-1"`
		String      string        `name:"string"      type:"string"      usage:"hello" value:"some string"`
		StringSlice []string      `name:"stringslice" type:"stringslice" usage:"hello" value:"some,string,slice"`
		Duration    time.Duration `name:"duration"    type:"duration"    usage:"hello" value:"2h1m10s"`
		//Custom      custom        `name:"custom"      type:"generic"     usage:"hello"`
	}

	sample := &Sample{}

	flags, err := FlagsFromStruct(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	ch := make(chan *cli.Context, 1)
	defer close(ch)

	app := cli.NewApp()
	app.Flags = flags
	app.Action = func(context *cli.Context) error {
		ch <- context
		return nil
	}

	err = app.Run(
		// XXX: First argument is a program name, so it is empty.
		[]string{
			"",
			"--bool",
			"--uint", "10",
			"--uint64", "10",
			"--int", "10",
			"--int64", "-10",
			"--float64", "10.05",
			"--intslice", "5", "--intslice", "4",
			"--int64slice", "5", "--int64slice", "4",
			"--string", "string some",
			"--stringslice", "and", "--stringslice", "others",
			"--duration", "1h",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}

	err = FlagsToStruct(<-ch, sample)
	if err != nil {
		t.Error(err)
		return
	}

	expectedSample := &Sample{
		Bool:        true,
		BoolT:       true,
		UInt:        10,
		UInt64:      10,
		Int:         10,
		Int64:       -10,
		Float64:     10.05,
		IntSlice:    []int{1, 2, 3, -1, 5, 4},
		Int64Slice:  []int64{1, 2, 3, -1, 5, 4},
		String:      "string some",
		StringSlice: []string{"some", "string", "slice", "and", "others"},
		Duration:    time.Hour,
	}

	assert.EqualValues(t, expectedSample, sample)
}
