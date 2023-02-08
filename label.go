package main

// isLabel will check if the line is a label and then add it to the symbol table.
func isLabel(line string) bool {
	if line[0] == '(' {
		return true
	}
	return false
}

// getLabelIndex will get the last index of the label.
func getLabelIndex(line string) int {
	for i := 0; i < len(line); i++ {
		if line[i] == ')' {
			return i
		}
	}
	return -1
}

// getLabelName will get the label name from the labelIndex.
func getLabelName(line string, labelIndex int) string {
	return line[1:labelIndex]
}
