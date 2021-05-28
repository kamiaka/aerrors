package aerrors

// Config for create *Err
type Config struct {
	priority    ErrorPriority
	formatError ErrorFormatter
	callerDepth int
	callerSkip  int
}

// DefaultConfig for create *Err.
var DefaultConfig = &Config{
	priority:    Error,
	formatError: NewFormatter("\n", ": "),
	callerDepth: 1,
	callerSkip:  0,
}

// Priority represents error priority.
func (c *Config) Priority() ErrorPriority {
	return c.priority
}

// WithPriority sets ErrorPriority and return receiver.
func (c *Config) WithPriority(p ErrorPriority) *Config {
	c.priority = p
	return c
}

// Formatter returns formater that format error messages.
// It is called by (*Err).FormatError.
func (c *Config) Formatter() ErrorFormatter {
	return c.formatError
}

// WithFormatter sets ErrorFormatter and return receiver.
func (c *Config) WithFormatter(f ErrorFormatter) *Config {
	c.formatError = f
	return c
}

// CallerDepth returns specified caller depth.
func (c *Config) CallerDepth() int {
	return c.callerDepth
}

// WithCallerDepth sets caller depth and return receiver.
func (c *Config) WithCallerDepth(n int) *Config {
	c.callerDepth = n
	return c
}

// CallerSkip returns specified caller skip count.
func (c *Config) CallerSkip() int {
	return c.callerSkip
}

// WithCallerSkip sets caller skip and return receiver.
func (c *Config) WithCallerSkip(n int) *Config {
	c.callerSkip = n
	return c
}

// Clone *Config.
func (c *Config) Clone() *Config {
	copy := *c
	return &copy
}

// Error returns new aerror's error from Config.
func (c *Config) Error(msg string, opts ...Option) *Err {
	return newErr(c, msg, opts...)
}

// Errorf formats according to a format specifier and returns the string as a value that satisfies error.
//
// If the format specifier has suffix `: %w` verb with an error operand, the returned error will implement an Unwrap method returning the operand.
func (c *Config) Errorf(format string, args ...interface{}) *Err {
	return errorf(c, format, args...)
}
