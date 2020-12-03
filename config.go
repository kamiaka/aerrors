package aerrors

// Config for create *Err
type Config struct {
	// Priority represents error priority.
	Priority ErrorPriority

	// FormatError formats error messages.
	// It is called by (*Err).FormatError.
	FormatError ErrorFormatter

	// Depth of callers.
	Depth int

	// Skip callers count.
	Skip int
}

// DefaultConfig for create *Err
var DefaultConfig = &Config{
	Priority:    Error,
	Depth:       1,
	Skip:        0,
	FormatError: NewFormatter("\n", ": "),
}

// Clone *Config.
func (c *Config) Clone() *Config {
	copy := *c
	return &copy
}
