package aerrors

// Option for create error `*Err`.
type Option func(*Config) *Config

// Priority option configures error priority.
func Priority(p ErrorPriority) Option {
	return func(c *Config) *Config {
		c.Priority = p
		return c
	}
}

// CallerDepth option configures callers depth.
func CallerDepth(n int) Option {
	return func(c *Config) *Config {
		c.CallerDepth = n
		return c
	}
}

// CallerSkip option configures callers skip.
func CallerSkip(n int) Option {
	return func(c *Config) *Config {
		c.CallerSkip = n
		return c
	}
}

// Formatter option configures error formatter.
func Formatter(f ErrorFormatter) Option {
	return func(c *Config) *Config {
		c.FormatError = f
		return c
	}
}
