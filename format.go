package aerrors

import (
	"golang.org/x/xerrors"
)

// ErrorFormatter is func for format error.
type ErrorFormatter func(xerrors.Printer, *Err) (next error)

// NewFormatter returns
func NewFormatter(sep, labelSep string) ErrorFormatter {
	return func(p xerrors.Printer, e *Err) (next error) {
		p.Print(e.msg)
		if p.Detail() {
			p.Print(sep, "priority", labelSep, e.priority)
			parent := e.parent
			for {
				if parent == nil {
					break
				}
				p.Print(sep, "parent", labelSep, parent.msg)
				parent = parent.parent
			}
			p.Print(sep, "callers", labelSep, e.callers.String())
			for _, v := range e.values {
				p.Print(sep, v.Label, labelSep, v.Value)
			}
			return e.wrappedError
		}
		return nil
	}
}
