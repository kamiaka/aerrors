package aerrors

import "fmt"

// ErrorPriority is error priority.
// Smaller value is higher priority.
type ErrorPriority int

// Built in priorities.
const (
	Emergency ErrorPriority = iota
	Alert
	Critical
	Error
	Warning
	Notice
	Info
	Debug
)

// PriorityNames for (ErrorPriority).Name
//
// It can overwrite for user defined priority.
//
//   aerrors.PriorityNames = map[aerrors.ErrorPriority]string{
//      Foo: "foo error",
//   }
var PriorityNames = map[ErrorPriority]string{
	Emergency: "Emergency",
	Alert:     "Alert",
	Critical:  "Critical",
	Error:     "Error",
	Warning:   "Warning",
	Notice:    "Notice",
	Info:      "Info",
	Debug:     "Debug",
}

// HigherThan reports whether the priority s is higher priority than t.
func (p ErrorPriority) HigherThan(q ErrorPriority) bool {
	return p < q
}

func (p ErrorPriority) String() string {
	if name, ok := PriorityNames[p]; ok {
		return name
	}
	return fmt.Sprintf("ErrorPriority(%d)", int(p))
}
