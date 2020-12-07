package aerrors

// Config for create *Err
type Config struct {
	// Priority represents error priority.
	Priority ErrorPriority

	// FormatError formats error messages.
	// It is called by (*Err).FormatError.
	FormatError ErrorFormatter

	// CallerDepth .
	CallerDepth int

	// CallerSkip.
	CallerSkip int
}

// DefaultConfig for create *Err
var DefaultConfig = &Config{
	Priority:    Error,
	FormatError: NewFormatter("\n", ": "),
	CallerDepth: 1,
	CallerSkip:  0,
}

// Clone *Config.
func (c *Config) Clone() *Config {
	copy := *c
	return &copy
}
