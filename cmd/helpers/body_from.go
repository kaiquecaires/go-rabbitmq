package helpers

import "strings"

func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
