package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Parser struct {
	lines []string
}

func NewParser(lines []string) *Parser {
	return &Parser{
		lines: lines,
	}
}

func (p *Parser) Parse() []string {
	// Remove all comments and empty lines from the lines.
	lines := removeAllCommentsAndEmptyLines(p.lines)

	// Set a label counter to keep track of the labels.
	labelCounter := 0

	// Set a symbol counter to keep track of the symbols.
	symbolCounter := 16

	// First loop to find all labels.
	for index, line := range lines {
		line = removeComments(line)
		line = removeSpaces(line)
		if !isEmptyLine(line) && !isComment(line) {
			if isLabel(line) {
				labelIndex := getLabelIndex(line)
				labelName := getLabelName(line, labelIndex)
				insertSymbol(labelName, index-labelCounter)
				labelCounter++
			}
		}
	}

	var parsedLines []string
	for _, line := range p.lines {
		if !isComment(line) && !isEmptyLine(line) && !isLabel(line) {
			line = removeComments(line)
			line = removeSpaces(line)
			if isSymbol(line) {
				if !isSymbolInTable(line[1:]) && !unicode.IsDigit(rune(line[1])) {
					insertSymbol(line[1:], symbolCounter)
					fmt.Println("Is added: ", line[1:], symbolCounter)
					symbolCounter++
				}
			}
			if isAddress(line) {
				parsedLines = append(parsedLines, parseAddress(line))
			} else {
				parsedLines = append(parsedLines, parseCommand(line))
			}
		}
	}
	return parsedLines
}

// removeAllCommentsAndEmptyLines removes all comments and empty lines.
func removeAllCommentsAndEmptyLines(lines []string) []string {
	var newLines []string
	for _, line := range lines {
		if !isComment(line) && !isEmptyLine(line) {
			newLines = append(newLines, line)
		}
	}
	return newLines
}

// removeComments will remove comments from the line.
func removeComments(line string) string {
	for i := 0; i < len(line); i++ {
		if line[i] == '/' {
			return line[:i]
		}
	}
	return line
}

// removeSpaces will remove spaces from the line.
func removeSpaces(line string) string {
	return strings.Replace(line, " ", "", -1)
}

// isComment will check if the line is a comment.
func isComment(line string) bool {
	if len(line) == 0 {
		return false
	}
	if line[0] == '/' {
		return true
	}
	return false
}

// isEmptyLine will check if the line is empty.
func isEmptyLine(line string) bool {
	if len(line) == 0 {
		return true
	}
	return false
}

// isCommand will check if the line is a command.
func isCommand(line string) bool {
	if line[0] == '@' || line[0] == '(' || line[0] == '/' || line[0] == ' ' {
		return false
	}
	return true
}

// isAddress will check if the line is an address.
func isAddress(line string) bool {
	if line[0] == '@' {
		return true
	}
	return false
}

// parseLine to binary.
func parseLine(line string) string {
	line = removeComments(line)
	line = removeSpaces(line)
	if isAddress(line) {
		return parseAddress(line)
	}
	return parseCommand(line)
}

// parseAddress to binary.
func parseAddress(line string) string {
	symbol := line[1:]
	if unicode.IsDigit(rune(symbol[0])) {
		return "0" + parseValue(symbol)
	}
	value := getSymbol(symbol)
	stringValue := strconv.Itoa(value)
	return "0" + parseValue(stringValue)
}

// parseCommand to binary.
func parseCommand(line string) string {
	comp, dest, jump := splitCommand(line)
	return "111" + comp + dest + jump
}

// splitCommand will split the command into dest, comp, jump.
func splitCommand(line string) (comp string, dest string, jump string) {
	line = removeComments(line)
	dest = ""
	comp = line
	jump = ""

	containsDest := strings.Contains(line, "=")
	containsJump := strings.Contains(line, ";")

	if containsDest {
		split := strings.Split(line, "=")
		dest = split[0]
		comp = split[1]
	}

	if containsJump {
		split := strings.Split(line, ";")
		jump = split[1]
		comp = split[0]
	}

	dest = getDest(dest)

	jump = getJump(jump)

	comp = getComp(comp)

	return

}

// parseValue to binary.
func parseValue(line string) string {
	return toBinary(line)
}

// getValue from the line.
func getValue(line string) string {
	lineWithoutComments := removeComments(line)
	return lineWithoutComments[1:]
}

// toBinary will parse the value to binary
func toBinary(value string) string {
	val, err := strconv.Atoi(value)
	if err != nil {
		log.Panicln(err)
	}

	// parse the int to 15 bit binary
	return fmt.Sprintf("%015b", val)
}
