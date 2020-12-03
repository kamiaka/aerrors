package aerrors

import (
	"reflect"
	"regexp"
	"testing"
	"time"
)

type fakeStringer struct {
	value string
}

func (s *fakeStringer) String() string {
	return s.value
}

func TestValue_generators(t *testing.T) {
	tmpDefaultStackDepth := DefaultStackDepth
	DefaultStackDepth = 2

	cases := []struct {
		value *Value
		want  *Value
	}{
		{
			value: String("string", "string value"),
			want: &Value{
				Label: "string",
				Value: "string value",
			},
		},
		{
			value: Stringer("stringer", &fakeStringer{value: "stringer value"}),
			want: &Value{
				Label: "stringer",
				Value: "stringer value",
			},
		},
		{
			value: Bool("bool", true),
			want: &Value{
				Label: "bool",
				Value: "true",
			},
		},
		{
			value: Bool("bool", false),
			want: &Value{
				Label: "bool",
				Value: "false",
			},
		},
		{
			value: Bytes("bytes", []byte("foo")),
			want: &Value{
				Label: "bytes",
				Value: "0x666f6f",
			},
		},
		{
			value: Byte("bytes", 'f'),
			want: &Value{
				Label: "bytes",
				Value: "0x66",
			},
		},
		{
			value: func() *Value {
				return Stack()
			}(),
			want: &Value{
				Label: "stack",
				Value: "aerrors.TestValue_generators.func1:github.com/kamiaka/aerrors/value_test.go:00, aerrors.TestValue_generators:github.com/kamiaka/aerrors/value_test.go:00",
			},
		},
		{
			value: func() *Value {
				return StackN(1)
			}(),
			want: &Value{
				Label: "stack",
				Value: "aerrors.TestValue_generators.func2:github.com/kamiaka/aerrors/value_test.go:00",
			},
		},
		{
			value: Int("int", 42),
			want: &Value{
				Label: "int",
				Value: "42",
			},
		},
		{
			value: Time("time", time.Date(2001, time.February, 3, 4, 5, 6, 7, time.FixedZone("+9000", 9*3600))),
			want: &Value{
				Label: "time",
				Value: "2001-02-03T04:05:06.000000007+09:00",
			},
		},
		{
			value: UTCTime("time", time.Date(2001, time.February, 3, 4, 5, 6, 7, time.FixedZone("+9000", 9*3600))),
			want: &Value{
				Label: "time",
				Value: "2001-02-02T19:05:06.000000007Z",
			},
		},
	}
	DefaultStackDepth = tmpDefaultStackDepth

	for i, tc := range cases {
		if !reflect.DeepEqual(tc.value, tc.want) {
			if tc.value.Label == "stack" && tc.value.Label == tc.want.Label && trimStackLine(tc.value.Value) == trimStackLine(tc.want.Value) {
				continue
			}
			t.Errorf("#%d:\nvalue: %#v\nwant:  %#v", i, tc.value, tc.want)
		}
	}
}

func trimStackLine(stack string) string {
	re := regexp.MustCompile(`:\d+`)
	return re.ReplaceAllString(stack, "")
}

func TestValue_WithLabel(t *testing.T) {
	cases := []struct {
		value *Value
		want  *Value
	}{
		{
			value: String("label", "value"),
			want:  &Value{Label: "label", Value: "value"},
		},
		{
			value: String("label", "value").WithLabel("new label"),
			want:  &Value{Label: "new label", Value: "value"},
		},
	}

	for i, tc := range cases {
		if !reflect.DeepEqual(tc.value, tc.want) {
			t.Errorf("#%d:\nvalue: %#v\nwant:  %#v", i, tc.value, tc.want)
		}
	}
}

func TestValue_String(t *testing.T) {
	cases := []struct {
		value *Value
		want  string
	}{
		{
			value: &Value{
				Label: "label",
				Value: "value",
			},
			want: "label: value",
		},
	}

	for i, tc := range cases {
		got := tc.value.String()
		if tc.want != got {
			t.Errorf("#%d: *Value.String()\ngot:  %#v\nwant: %#v", i, got, tc.want)
		}
	}
}
