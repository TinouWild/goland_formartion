package basics

import "strings"

func ToLoverStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}