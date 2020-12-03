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
			p.Print(sep, "priority", labelSep, e.config.Priority)
			if e.parent != nil {
				p.Print(sep, "parent", labelSep, e.parent.msg)
			}
			if e.wrappedError != nil {
				p.Print(sep, "origin", labelSep, e.wrappedError.Error())
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
