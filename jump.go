package main

import "strings"

var jump = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

// getJump will get the binary code for the jump.
func getJump(value string) string {
	return jump[value]
}

// Check if the line contains a jump by checking if it contains a ';'.
// If it does, then we need to get the jump.
// If it doesn't, then we need to return an empty string.
func getJumpBinary(line string) string {
	jump := ""
	if strings.Contains(line, ";") {
		return line[strings.Index(line, ";")+1:]
	}
	jump = getJump(jump)
	return jump
}
