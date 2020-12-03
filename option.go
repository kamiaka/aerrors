package aerrors

// Option for create error `*Err`.
type Option func(*Config)

// Priority option configures error priority.
func Priority(p ErrorPriority) Option {
	return func(c *Config) {
		c.Priority = p
	}
}

// Depth option configures callers depth.
func Depth(n int) Option {
	return func(c *Config) {
		c.Depth = n
	}
}

// Skip option configures callers skip.
func Skip(n int) Option {
	return func(c *Config) {
		c.Skip = n
	}
}

// Formatter option configures error formatter.
func Formatter(f ErrorFormatter) Option {
	return func(c *Config) {
		c.FormatError = f
	}
}
