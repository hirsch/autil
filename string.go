package autil

import "strings"

// StringBetween returns the strings found between two delimiters.
func StringBetween(str, start, end string) (between []string) {
	between = make([]string, 0)
	fsplit := strings.Split(str, start)
	for i := 1; i < len(fsplit); i++ {
		ssplit := strings.SplitN(fsplit[i], end, 2)
		if len(ssplit) >= 2 {
			between = append(between, ssplit[0])
		}
	}
	return
}
