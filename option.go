package aerrors

// Option for create error `*Err`.
type Option func(*Config) *Config

// Priority option configures error priority.
func Priority(p ErrorPriority) Option {
	return func(c *Config) *Config {
		return c.WithPriority(p)
	}
}

// CallerDepth option configures callers depth.
func CallerDepth(n int) Option {
	return func(c *Config) *Config {
		return c.WithCallerDepth(n)
	}
}

// CallerSkip option configures callers skip.
func CallerSkip(n int) Option {
	return func(c *Config) *Config {
		return c.WithCallerSkip(n)
	}
}

// Formatter option configures error formatter.
func Formatter(f ErrorFormatter) Option {
	return func(c *Config) *Config {
		return c.WithFormatter(f)
	}
}
