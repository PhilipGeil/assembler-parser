package main

import "strings"

var dest = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
}

// getDest will get the binary code for the dest.
func getDest(value string) string {
	return dest[value]
}

// Check if the line contains a dest by checking if it contains an '='.
// If it does, then we need to get the dest.
// If it doesn't, then we need to return an empty string.
func getDestBinary(line string) string {
	dest := ""
	if strings.Contains(line, "=") {
		return line[:strings.Index(line, "=")]
	}
	dest = getDest(dest)
	return dest
}
