package aerrors

import "errors"

// AsErr returns casted `*Err` error and whether cas succeeded.
//
// It is the same as the code below.
//   var e *aerrors.Err
//   ok := errors.As(&e)
func AsErr(err error) (e *Err, ok bool) {
	if errors.As(err, &e) {
		return e, true
	}
	return nil, false
}
