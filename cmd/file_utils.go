package main

import "strings"

func filenameParts(filename string, separator string) []string {
	parts := strings.Split(filename, separator)

	return parts
}
