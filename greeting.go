package main

import (
	"strings"
)

func getInitials(name string) (val1 string, val2 string) {
	str := strings.ToUpper(name)
	names := strings.Split(str, " ")
	var initialRange []string
	for _, value := range names {
		initialRange = append(initialRange, value[:1])
	}
	if len(initialRange) > 1 {
		return initialRange[0], initialRange[1]
	}
	return initialRange[0], "_"

}
