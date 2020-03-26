package convert_test

import (
	"testing"

	"time"

	"github.com/Eun/go-convert/internal/testhelpers"
)

func TestUint8(t *testing.T) {
	tests := []testhelpers.TestCase{
		// nil
		{nil, uint8(0), uint8(0), `unable to convert convert.NilValue to uint8: no recipe`, nil},
		// string
		{"6", uint8(0), uint8(6), "", nil},
		{"", uint8(0), uint8(0), "", nil},
		{"Hello World", uint8(0), uint8(0), `unable to convert string to uint8: strconv.ParseUint: parsing "Hello World": invalid syntax`, nil},
		// bool
		{true, uint8(0), uint8(1), "", nil},
		{false, uint8(0), uint8(0), "", nil},
		// int
		{6, uint8(0), uint8(6), "", nil},
		// int8
		{int8(6), uint8(0), uint8(6), "", nil},
		// int16
		{int16(6), uint8(0), uint8(6), "", nil},
		// int32
		{int32(6), uint8(0), uint8(6), "", nil},
		// int64
		{int64(6), uint8(0), uint8(6), "", nil},
		// uint
		{uint(6), uint8(0), uint8(6), "", nil},
		// int8
		{uint8(6), uint8(0), uint8(6), "", nil},
		// uint16
		{uint16(6), uint8(0), uint8(6), "", nil},
		// uint32
		{uint32(6), uint8(0), uint8(6), "", nil},
		// uint64
		{uint64(6), uint8(0), uint8(6), "", nil},
		// float32
		{float32(6), uint8(0), uint8(6), "", nil},
		// float64
		{float64(6), uint8(0), uint8(6), "", nil},
		// slice
		{[]int{'H', 'e', 'l', 'l', 'o'}, uint8(0), uint8(0), "unable to convert []int to uint8: no recipe", nil},
		{[]byte{'H', 'e', 'l', 'l', 'o'}, uint8(0), uint8(0), "unable to convert []uint8 to uint8: no recipe", nil},
		{[]rune{'H', 'e', 'l', 'l', 'o'}, uint8(0), uint8(0), "unable to convert []int32 to uint8: no recipe", nil},
		{[]string{"H", "e", "l", "l", "o"}, uint8(0), uint8(0), "unable to convert []string to uint8: no recipe", nil},
		// struct
		{struct{}{}, uint8(0), uint8(0), "unable to convert struct {} to uint8: no recipe", nil},
		// time
		{time.Unix(10, 10), uint8(10), uint8(10), "", nil},
	}

	for i, test := range tests {
		testhelpers.RunTest(t, test, i)
	}
}
