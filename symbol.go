package main

var symbols = map[string]int{
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
}

// insertSymbol will insert a symbol into the symbol table.
func insertSymbol(symbol string, address int) {
	symbols[symbol] = address
}

// getSymbol will get the address of a symbol from the symbol table.
func getSymbol(symbol string) int {
	return symbols[symbol]
}

// isSymbol will check if the line is a symbol.
func isSymbol(line string) bool {
	if line[0] == '@' {
		return true
	}
	return false
}

// getSymbolName will get the symbol name from the line.
func getSymbolName(line string) string {
	lineWithoutComments := removeComments(line)
	return lineWithoutComments[1:]
}

// Check if the symbol is in the symbol table.
func isSymbolInTable(symbol string) bool {
	_, ok := symbols[symbol]
	return ok
}
