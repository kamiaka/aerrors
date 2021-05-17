package stack

import "strings"

// simpleFunc returns path removed function name.
//   e.g.,
//     given:  path/to/pkgdir.(Type.)Method
//     return: pkgdir.(Type.)Method
func simpleFunc(funcName string) string {
	if index := strings.LastIndex(funcName, "/"); index >= 0 {
		return funcName[index+1:]
	}
	return funcName
}
