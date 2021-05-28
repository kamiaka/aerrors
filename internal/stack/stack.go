package stack

import (
	"bytes"
	"runtime"
	"strconv"
)

// Frames ...
type Frames struct {
	*runtime.Frames
}

// Callers ...
func Callers(depth, skip int) *Frames {
	pc := make([]uintptr, depth)

	runtime.Callers(skip+2, pc)

	return &Frames{
		runtime.CallersFrames(pc),
	}
}

// Format frames to string by specified separators.
func (f *Frames) Format(sep, funcSep, lineSep string) string {
	var b bytes.Buffer
	inMore := false
	for {
		frame, more := f.Next()
		if frame.Function == "runtime.main" {
			break
		}
		if inMore {
			b.WriteString(sep)
		}
		b.WriteString(simpleFunc(frame.Function) + funcSep + frame.File + lineSep + strconv.Itoa(frame.Line))
		if more {
			inMore = true
			continue
		}
		break
	}

	return b.String()
}

// String returns formatted string.
//   e.g., pkg(.Type).Func:path/to/file.go:line, ...
func (f *Frames) String() string {
	return f.Format(", ", ":", ":")
}
