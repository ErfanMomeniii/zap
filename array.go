// Copyright (c) 2016 Uber Technologies, Inc.
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

package zap

import (
	"time"

	"go.uber.org/zap/zapcore"
)

// Array constructs a field with the given key and ArrayMarshaler. It provides
// a flexible, but still type-safe and efficient, way to add array-like types
// to the logging context. The struct's MarshalLogArray method is called lazily.
func Array(key string, val zapcore.ArrayMarshaler) zapcore.Field {
	return zapcore.Field{Key: key, Type: zapcore.ArrayMarshalerType, Interface: val}
}

// Bools constructs a field that carries a slice of bools.
func Bools(key string, bs []bool) zapcore.Field {
	return Array(key, bools(bs))
}

// Bytes constructs a field that carries a slice of bytes. Note that most
// encoders represent byte slices as arrays of integers, not strings.
func Bytes(key string, bs []byte) zapcore.Field {
	return Array(key, bytes(bs))
}

// Complex128s constructs a field that carries a slice of complex numbers.
func Complex128s(key string, nums []complex128) zapcore.Field {
	return Array(key, complex128s(nums))
}

// Complex64s constructs a field that carries a slice of complex numbers.
func Complex64s(key string, nums []complex64) zapcore.Field {
	return Array(key, complex64s(nums))
}

// Float64s constructs a field that carries a slice of floats.
func Float64s(key string, nums []float64) zapcore.Field {
	return Array(key, float64s(nums))
}

// Float32s constructs a field that carries a slice of floats.
func Float32s(key string, nums []float32) zapcore.Field {
	return Array(key, float32s(nums))
}

// Ints constructs a field that carries a slice of integers.
func Ints(key string, nums []int) zapcore.Field {
	return Array(key, ints(nums))
}

// Int64s constructs a field that carries a slice of integers.
func Int64s(key string, nums []int64) zapcore.Field {
	return Array(key, int64s(nums))
}

// Int32s constructs a field that carries a slice of integers.
func Int32s(key string, nums []int32) zapcore.Field {
	return Array(key, int32s(nums))
}

// Int16s constructs a field that carries a slice of integers.
func Int16s(key string, nums []int16) zapcore.Field {
	return Array(key, int16s(nums))
}

// Int8s constructs a field that carries a slice of integers.
func Int8s(key string, nums []int8) zapcore.Field {
	return Array(key, int8s(nums))
}

// Runes constructs a field that carries a slice of runes. Note that most
// encoders will represent a slice of runes as an array of integers, not arrays
// of characters or Unicode code points.
func Runes(key string, rs []rune) zapcore.Field {
	return Array(key, runes(rs))
}

// Strings constructs a field that carries a slice of strings.
func Strings(key string, ss []string) zapcore.Field {
	return Array(key, stringArray(ss))
}

// Uints constructs a field that carries a slice of unsigned integers.
func Uints(key string, nums []uint) zapcore.Field {
	return Array(key, uints(nums))
}

// Uint64s constructs a field that carries a slice of unsigned integers.
func Uint64s(key string, nums []uint64) zapcore.Field {
	return Array(key, uint64s(nums))
}

// Uint32s constructs a field that carries a slice of unsigned integers.
func Uint32s(key string, nums []uint32) zapcore.Field {
	return Array(key, uint32s(nums))
}

// Uint16s constructs a field that carries a slice of unsigned integers.
func Uint16s(key string, nums []uint16) zapcore.Field {
	return Array(key, uint16s(nums))
}

// Uint8s constructs a field that carries a slice of unsigned integers.
func Uint8s(key string, nums []uint8) zapcore.Field {
	return Array(key, uint8s(nums))
}

// Uintptrs constructs a field that carries a slice of pointer addresses.
func Uintptrs(key string, us []uintptr) zapcore.Field {
	return Array(key, uintptrs(us))
}

// Errors constructs a field that carries a slice of errors.
func Errors(key string, errs []error) zapcore.Field {
	return Array(key, errArray(errs))
}

// Durations constructs a field that carries a slice of time.Durations. Like
// Duration, it represents each duration as an integer number of nanoseconds.
func Durations(key string, ds []time.Duration) zapcore.Field {
	return Array(key, durations(ds))
}

type bools []bool

func (bs bools) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range bs {
		arr.AppendBool(bs[i])
	}
	return nil
}

type bytes []byte

func (bs bytes) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range bs {
		arr.AppendByte(bs[i])
	}
	return nil
}

type complex128s []complex128

func (nums complex128s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendComplex128(nums[i])
	}
	return nil
}

type complex64s []complex64

func (nums complex64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendComplex64(nums[i])
	}
	return nil
}

type float64s []float64

func (nums float64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendFloat64(nums[i])
	}
	return nil
}

type float32s []float32

func (nums float32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendFloat32(nums[i])
	}
	return nil
}

type ints []int

func (nums ints) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendInt(nums[i])
	}
	return nil
}

type int64s []int64

func (nums int64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendInt64(nums[i])
	}
	return nil
}

type int32s []int32

func (nums int32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendInt32(nums[i])
	}
	return nil
}

type int16s []int16

func (nums int16s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendInt16(nums[i])
	}
	return nil
}

type int8s []int8

func (nums int8s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendInt8(nums[i])
	}
	return nil
}

type runes []rune

func (rs runes) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range rs {
		arr.AppendRune(rs[i])
	}
	return nil
}

type stringArray []string

func (ss stringArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range ss {
		arr.AppendString(ss[i])
	}
	return nil
}

type uints []uint

func (nums uints) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendUint(nums[i])
	}
	return nil
}

type uint64s []uint64

func (nums uint64s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendUint64(nums[i])
	}
	return nil
}

type uint32s []uint32

func (nums uint32s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendUint32(nums[i])
	}
	return nil
}

type uint16s []uint16

func (nums uint16s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendUint16(nums[i])
	}
	return nil
}

type uint8s []uint8

func (nums uint8s) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendUint8(nums[i])
	}
	return nil
}

type uintptrs []uintptr

func (nums uintptrs) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range nums {
		arr.AppendUintptr(nums[i])
	}
	return nil
}

type errArray []error

func (errs errArray) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range errs {
		if errs[i] == nil {
			continue
		}
		arr.AppendString(errs[i].Error())
	}
	return nil
}

type durations []time.Duration

func (ds durations) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for i := range ds {
		arr.AppendInt64(int64(ds[i]))
	}
	return nil
}