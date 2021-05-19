package aerrors

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kamiaka/aerrors/internal/stack"
)

// Value is labeled value.
//
// Using `(*Err).With(values...)`
type Value struct {
	Label string
	Value string
}

// WithLabel sets the label `l` and returns it receiver.
func (v *Value) WithLabel(l string) *Value {
	v.Label = l
	return v
}

// String returns label and value string.
func (v *Value) String() string {
	return fmt.Sprintf("%s: %s", v.Label, v.Value)
}

// String returns Value.
func String(l, v string) *Value {
	return &Value{
		Label: l,
		Value: v,
	}
}

// Stringer returns Value.
func Stringer(l string, v interface{ String() string }) *Value {
	return &Value{
		Label: l,
		Value: v.String(),
	}
}

// Stringf returns formatted string Value.
func Stringf(l string, v string, args ...interface{}) *Value {
	return &Value{
		Label: l,
		Value: fmt.Sprintf(v, args...),
	}
}

// Bool returns Value.
func Bool(l string, v bool) *Value {
	var s string
	if v {
		s = "true"
	} else {
		s = "false"
	}
	return &Value{
		Label: l,
		Value: s,
	}
}

var hexPrefix = []byte("0x")
var digits = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

// Bytes returns Value.
func Bytes(l string, v []byte) *Value {
	buf := hexPrefix
	for _, b := range v {
		buf = append(buf, digits[b/16], digits[b%16])
	}
	return &Value{
		Label: l,
		Value: string(buf),
	}
}

// Byte returns Value.
func Byte(l string, b byte) *Value {
	return &Value{
		Label: l,
		Value: string(append(hexPrefix, digits[b/16], digits[b%16])),
	}
}

// Int returns Value.
func Int(l string, v int) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int8 returns Value.
func Int8(l string, v int8) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int16 returns Value.
func Int16(l string, v int16) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int32 returns Value.
func Int32(l string, v int32) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int64 returns Value.
func Int64(l string, v int64) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(v, 10),
	}
}

// Uint returns Value.
func Uint(l string, v uint) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint8 returns Value.
func Uint8(l string, v uint8) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint16 returns Value.
func Uint16(l string, v uint16) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint32 returns Value.
func Uint32(l string, v uint32) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint64 returns Value.
func Uint64(l string, v uint64) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Float32 returns Value of float32.
func Float32(l string, v float32) *Value {
	return &Value{
		Label: l,
		Value: fmt.Sprint(v),
	}
}

// Float64 returns Value of float64.
func Float64(l string, v float64) *Value {
	return &Value{
		Label: l,
		Value: fmt.Sprint(v),
	}
}

// Time returns Value of time.
func Time(l string, v time.Time) *Value {
	return &Value{
		Label: l,
		Value: v.Format(time.RFC3339Nano),
	}
}

// UTCTime returns Value of UTC time.
func UTCTime(l string, v time.Time) *Value {
	return Time(l, v.UTC())
}

// DefaultStackDepth is the depth used when call Stack.
var DefaultStackDepth = 16

// Stack returns Value of stack trace.
// depth is determined by DefaultStackDepth.
func Stack(skip int) *Value {
	return &Value{
		Label: "stack",
		Value: stack.Callers(DefaultStackDepth, skip+1).String(),
	}
}

// StackN returns Value of stack trace for N layers.
func StackN(depth, skip int) *Value {
	return &Value{
		Label: "stack",
		Value: stack.Callers(depth, skip+1).String(),
	}
}
