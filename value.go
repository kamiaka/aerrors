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

type Values []*Value

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

// String appends string Value and return Values.
func (ls Values) String(l, v string) Values {
	return append(ls, String(l, v))
}

// Stringer returns Value.
func Stringer(l string, v interface{ String() string }) *Value {
	return &Value{
		Label: l,
		Value: v.String(),
	}
}

// Stringer appends stringer Value and return Values.
func (ls Values) Stringer(l string, v interface{ String() string }) Values {
	return append(ls, Stringer(l, v))
}

// Stringf returns formatted string Value.
func Stringf(l string, format string, args ...interface{}) *Value {
	return &Value{
		Label: l,
		Value: fmt.Sprintf(format, args...),
	}
}

// Stringf appends formatted string Value and return Values.
func (ls Values) Stringf(l string, format string, args ...interface{}) Values {
	return append(ls, Stringf(l, format, args...))
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

// Bool appends bool Value and return Values.
func (ls Values) Bool(l string, v bool) Values {
	return append(ls, Bool(l, v))
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

// Bytes appends bytes Value and return Values.
func (ls Values) Bytes(l string, v []byte) Values {
	return append(ls, Bytes(l, v))
}

// Byte returns Value.
func Byte(l string, b byte) *Value {
	return &Value{
		Label: l,
		Value: string(append(hexPrefix, digits[b/16], digits[b%16])),
	}
}

// Byte appends byte Value and return Values.
func (ls Values) Byte(l string, v byte) Values {
	return append(ls, Byte(l, v))
}

// Int returns Value.
func Int(l string, v int) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int appends int Value and return Values.
func (ls Values) Int(l string, v int) Values {
	return append(ls, Int(l, v))
}

// Int8 returns Value.
func Int8(l string, v int8) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int8 appends Int8 Value and return Values.
func (ls Values) Int8(l string, v int8) Values {
	return append(ls, Int8(l, v))
}

// Int16 returns Value.
func Int16(l string, v int16) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int16 appends Int16 Value and return Values.
func (ls Values) Int16(l string, v int16) Values {
	return append(ls, Int16(l, v))
}

// Int32 returns Value.
func Int32(l string, v int32) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(int64(v), 10),
	}
}

// Int32 appends Int32 Value and return Values.
func (ls Values) Int32(l string, v int32) Values {
	return append(ls, Int32(l, v))
}

// Int64 returns Value.
func Int64(l string, v int64) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatInt(v, 10),
	}
}

// Int64 appends Int64 Value and return Values.
func (ls Values) Int64(l string, v int64) Values {
	return append(ls, Int64(l, v))
}

// Uint returns Value.
func Uint(l string, v uint) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint appends Uint Value and return Values.
func (ls Values) Uint(l string, v uint) Values {
	return append(ls, Uint(l, v))
}

// Uint8 returns Value.
func Uint8(l string, v uint8) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint8 appends Uint8 Value and return Values.
func (ls Values) Uint8(l string, v uint8) Values {
	return append(ls, Uint8(l, v))
}

// Uint16 returns Value.
func Uint16(l string, v uint16) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint16 appends Uint16 Value and return Values.
func (ls Values) Uint16(l string, v uint16) Values {
	return append(ls, Uint16(l, v))
}

// Uint32 returns Value.
func Uint32(l string, v uint32) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint32 appends Uint32 Value and return Values.
func (ls Values) Uint32(l string, v uint32) Values {
	return append(ls, Uint32(l, v))
}

// Uint64 returns Value.
func Uint64(l string, v uint64) *Value {
	return &Value{
		Label: l,
		Value: strconv.FormatUint(uint64(v), 10),
	}
}

// Uint64 appends Uint64 Value and return Values.
func (ls Values) Uint64(l string, v uint64) Values {
	return append(ls, Uint64(l, v))
}

// Float32 returns Value of float32.
func Float32(l string, v float32) *Value {
	return &Value{
		Label: l,
		Value: fmt.Sprint(v),
	}
}

// Float32 appends Float32 Value and return Values.
func (ls Values) Float32(l string, v float32) Values {
	return append(ls, Float32(l, v))
}

// Float64 returns Value of float64.
func Float64(l string, v float64) *Value {
	return &Value{
		Label: l,
		Value: fmt.Sprint(v),
	}
}

// Float64 appends Float64 Value and return Values.
func (ls Values) Float64(l string, v float64) Values {
	return append(ls, Float64(l, v))
}

// Time returns Value of time.
func Time(l string, v time.Time) *Value {
	return &Value{
		Label: l,
		Value: v.Format(time.RFC3339Nano),
	}
}

// Time appends Time Value and return Values.
func (ls Values) Time(l string, v time.Time) Values {
	return append(ls, Time(l, v))
}

// UTCTime returns Value of UTC time.
func UTCTime(l string, v time.Time) *Value {
	return Time(l, v.UTC())
}

// UTCTime appends UTCTime Value and return Values.
func (ls Values) UTCTime(l string, v time.Time) Values {
	return append(ls, UTCTime(l, v))
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

// Stack appends Stack Value and return Values.
func (ls Values) Stack(skip int) Values {
	return append(ls, Stack(skip))
}

// StackN returns Value of stack trace for N layers.
func StackN(depth, skip int) *Value {
	return &Value{
		Label: "stack",
		Value: stack.Callers(depth, skip+1).String(),
	}
}

// StackN appends Stack Value and return Values.
func (ls Values) StackN(depth, skip int) Values {
	return append(ls, StackN(depth, skip+1))
}
