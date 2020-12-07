package aerrors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/kamiaka/aerrors/internal/stack"
	"golang.org/x/xerrors"
)

// Err is aerror's error. It implements interface `error`.
type Err struct {
	config       *Config
	msg          string
	parent       *Err
	wrappedError error
	callers      *stack.Frames
	priority     ErrorPriority
	formatError  ErrorFormatter
	values       []*Value
}

// New aerror's error with options.
func New(msg string, opts ...Option) *Err {
	conf := DefaultConfig.Clone()
	for _, opt := range opts {
		conf = opt(conf)
	}

	return &Err{
		config:      conf,
		msg:         msg,
		callers:     stack.Callers(conf.CallerDepth, conf.CallerSkip+1),
		priority:    conf.Priority,
		formatError: conf.FormatError,
	}
}

// Errorf formats according to a format specifier and returns the string as a value that satisfies error.
//
// If the format specifier has suffix `: %w` verb with an error operand, the returned error will implement an Unwrap method returning the operand.
func Errorf(format string, args ...interface{}) *Err {
	format, wrappedError := wrappedFormat(format, args)
	conf := DefaultConfig.Clone()

	return &Err{
		config:       conf,
		msg:          fmt.Sprintf(format, args...),
		callers:      stack.Callers(conf.CallerDepth, conf.CallerSkip+1),
		priority:     conf.Priority,
		formatError:  conf.FormatError,
		wrappedError: wrappedError,
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

	conf := e.config.Clone()
	for _, opt := range opts {
		conf = opt(conf)
	}

	child.msg = msg
	child.callers = stack.Callers(conf.CallerDepth, conf.CallerSkip+2)
	child.parent = e
	child.priority = conf.Priority
	child.formatError = conf.FormatError

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

// Wrap error `err`.
func (e *Err) Wrap(err error) *Err {
	child := e.newChild(err.Error())
	child.wrappedError = err
	return child
}

// Unwrap error.
func (e *Err) Unwrap() error {
	return e.wrappedError
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

// Values set by `With` method.
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
func (e *Err) Config() *Config {
	return e.config
}
