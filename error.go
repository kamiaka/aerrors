package aerrors

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/kamiaka/aerrors/internal/stack"
	"golang.org/x/xerrors"
)

// Err is aerror's error. It implements interface `error`.
type Err struct {
	msg          string
	parent       *Err
	wrappedError error
	callers      *stack.Frames
	priority     ErrorPriority
	formatError  ErrorFormatter
	values       []*Value
	childConf    *Config
}

// New aerror's error with options.
func New(msg string, opts ...Option) *Err {
	return newErr(DefaultConfig, msg, opts...)
}

func newErr(conf *Config, msg string, opts ...Option) *Err {
	conf = conf.Clone()

	for _, opt := range opts {
		conf = opt(conf)
	}

	return &Err{
		msg:         msg,
		callers:     stack.Callers(conf.callerDepth, conf.callerSkip+2),
		priority:    conf.priority,
		formatError: conf.formatError,
		childConf:   conf.WithCallerSkip(0),
	}
}

// Errorf formats according to a format specifier and returns the string as a value that satisfies error.
//
// If the format specifier has suffix `: %w` verb with an error operand, the returned error will implement an Unwrap method returning the operand.
func Errorf(format string, args ...interface{}) *Err {
	return errorf(DefaultConfig, format, args...)
}

func errorf(conf *Config, format string, args ...interface{}) *Err {
	conf = conf.Clone()
	format, wrappedError := wrappedFormat(format, args)

	return &Err{
		msg:          fmt.Sprintf(format, args...),
		callers:      stack.Callers(conf.callerDepth, conf.callerSkip+2),
		priority:     conf.priority,
		formatError:  conf.formatError,
		wrappedError: wrappedError,
		childConf:    conf.WithCallerSkip(0),
	}
}

func wrappedFormat(format string, args []interface{}) (newFormat string, wrappedError error) {
	isWrapped := strings.HasSuffix(format, ": %w")
	if isWrapped && len(args) > 0 {
		format = string(append([]byte(format)[:len(format)-4])) + ": %v"
		if e, ok := args[len(args)-1].(error); ok {
			wrappedError = e
		}
	}
	return format, wrappedError
}

// Error implements interface `error`.
func (e *Err) Error() string {
	return e.msg
}

func (e *Err) clone() *Err {
	clone := *e
	return &clone
}

func (e *Err) newChild(msg string, opts ...Option) *Err {
	child := e.clone()

	conf := e.childConf.Clone().WithCallerSkip(0)
	for _, opt := range opts {
		conf = opt(conf)
	}

	child.msg = msg
	child.callers = stack.Callers(conf.callerDepth, conf.callerSkip+2)
	child.parent = e
	child.priority = conf.priority
	child.formatError = conf.formatError
	child.childConf = conf.WithCallerSkip(0)

	return child
}

// New child *Err.
func (e *Err) New(msg string, options ...Option) *Err {
	child := e.newChild(msg, options...)
	child.wrappedError = nil
	return child
}

// Errorf returns new child *Err with message.
func (e *Err) Errorf(format string, args ...interface{}) *Err {
	format, wrappedError := wrappedFormat(format, args)

	child := e.newChild(fmt.Sprintf(format, args...))
	child.wrappedError = wrappedError

	return child
}

// Wrap specified error `err`.
func (e *Err) Wrap(err error, opts ...Option) *Err {
	child := e.newChild(e.Error(), opts...)
	child.wrappedError = err
	return child
}

// Unwrap error.
func (e *Err) Unwrap() error {
	return e.wrappedError
}

// WithError sets wrapped error and returns receiver.
func (e *Err) WithError(err error) *Err {
	e.wrappedError = err
	return e
}

// Is reports whether the error `err` is `e`
func (e *Err) Is(err error) bool {
	return e == err || (e.parent != nil && e.parent.Is(err))
}

// Format implements interface `fmt.Formatter`
func (e *Err) Format(s fmt.State, verb rune) {
	xerrors.FormatError(e, s, verb)
}

// FormatError implements interface `xerrors.Formatter`
func (e *Err) FormatError(p xerrors.Printer) (next error) {
	e.formatError(p, e)
	if p.Detail() {
		return e.wrappedError
	}
	return nil
}

// Parent return parent *Err.
func (e *Err) Parent() *Err {
	return e.parent
}

// WithValue sets the `values` and returns receiver.
func (e *Err) WithValue(values ...*Value) *Err {
	e.values = append(e.values, values...)
	return e
}

// Values set by `WithValue` method.
func (e *Err) Values() []*Value {
	return e.values
}

// Callers returns error callers.
func (e *Err) Callers() *runtime.Frames {
	return e.callers.Frames
}

// Priority returns error priority.
func (e *Err) Priority() ErrorPriority {
	return e.priority
}

// WithPriority sets error priority and returns receiver.
func (e *Err) WithPriority(p ErrorPriority) *Err {
	e.priority = p
	return e
}

// Config returns aerror's config.
//
// Deprecated: Use *Err.ChildConfig.
func (e *Err) Config() *Config {
	return e.childConf
}

// ChildConfig returns *Config for new child.
func (e *Err) ChildConfig() *Config {
	return e.childConf
}

// WithChildConfig sets *Config for new child and receiver.
func (e *Err) WithChildConfig(c *Config) *Err {
	e.childConf = c
	return e
}

// WithString appends string Value and returns receiver.
func (e *Err) WithString(l, v string) *Err {
	e.values = append(e.values, String(l, v))
	return e
}

// WithStringer appends stringer Value and returns receiver.
func (e *Err) WithStringer(l string, v interface{ String() string }) *Err {
	e.values = append(e.values, Stringer(l, v))
	return e
}

// WithStringf appends formatted string Value and returns receiver.
func (e *Err) WithStringf(l string, format string, args ...interface{}) *Err {
	e.values = append(e.values, Stringf(l, format, args...))
	return e
}

// WithBool appends bool Value and returns receiver.
func (e *Err) WithBool(l string, v bool) *Err {
	e.values = append(e.values, Bool(l, v))
	return e
}

// WithBytes appends bytes Value and returns receiver.
func (e *Err) WithBytes(l string, v []byte) *Err {
	e.values = append(e.values, Bytes(l, v))
	return e
}

// WithByte appends byte Value and returns receiver.
func (e *Err) WithByte(l string, v byte) *Err {
	e.values = append(e.values, Byte(l, v))
	return e
}

// WithRune appends rune Value and returns receiver.
func (e *Err) WithRune(l string, v rune) *Err {
	e.values = append(e.values, Rune(l, v))
	return e
}

// WithInt appends int Value and returns receiver.
func (e *Err) WithInt(l string, v int) *Err {
	e.values = append(e.values, Int(l, v))
	return e
}

// WithInt8 appends Int8 Value and returns receiver.
func (e *Err) WithInt8(l string, v int8) *Err {
	e.values = append(e.values, Int8(l, v))
	return e
}

// WithInt16 appends Int16 Value and returns receiver.
func (e *Err) WithInt16(l string, v int16) *Err {
	e.values = append(e.values, Int16(l, v))
	return e
}

// WithInt32 appends Int32 Value and returns receiver.
func (e *Err) WithInt32(l string, v int32) *Err {
	e.values = append(e.values, Int32(l, v))
	return e
}

// WithInt64 appends Int64 Value and returns receiver.
func (e *Err) WithInt64(l string, v int64) *Err {
	e.values = append(e.values, Int64(l, v))
	return e
}

// WithUint appends Uint Value and returns receiver.
func (e *Err) WithUint(l string, v uint) *Err {
	e.values = append(e.values, Uint(l, v))
	return e
}

// WithUint8 appends Uint8 Value and returns receiver.
func (e *Err) WithUint8(l string, v uint8) *Err {
	e.values = append(e.values, Uint8(l, v))
	return e
}

// WithUint16 appends Uint16 Value and returns receiver.
func (e *Err) WithUint16(l string, v uint16) *Err {
	e.values = append(e.values, Uint16(l, v))
	return e
}

// WithUint32 appends Uint32 Value and returns receiver.
func (e *Err) WithUint32(l string, v uint32) *Err {
	e.values = append(e.values, Uint32(l, v))
	return e
}

// WithUint64 appends Uint64 Value and returns receiver.
func (e *Err) WithUint64(l string, v uint64) *Err {
	e.values = append(e.values, Uint64(l, v))
	return e
}

// WithFloat32 appends Float32 Value and returns receiver.
func (e *Err) WithFloat32(l string, v float32) *Err {
	e.values = append(e.values, Float32(l, v))
	return e
}

// WithFloat64 appends Float64 Value and returns receiver.
func (e *Err) WithFloat64(l string, v float64) *Err {
	e.values = append(e.values, Float64(l, v))
	return e
}

// WithTime appends Time Value and returns receiver.
func (e *Err) WithTime(l string, v time.Time) *Err {
	e.values = append(e.values, Time(l, v))
	return e
}

// WithUTCTime appends UTCTime Value and returns receiver.
func (e *Err) WithUTCTime(l string, v time.Time) *Err {
	e.values = append(e.values, UTCTime(l, v))
	return e
}

// WithStack appends Stack Value and returns receiver.
func (e *Err) WithStack(skip int) *Err {
	e.values = append(e.values, Stack(skip+1))
	return e
}

// WithStackN appends Stack Value and returns receiver.
func (e *Err) WithStackN(depth, skip int) *Err {
	e.values = append(e.values, StackN(depth, skip+1))
	return e
}
