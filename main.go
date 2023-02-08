package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("Pong.asm")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\r\n")
	// fmt.Println(len(lines))
	// fmt.Println(lines)
	parser := NewParser(lines)
	writeToFile(parser.Parse())

}

func writeToFile(lines []string) {
	// Open file using os.Create()
	file, err := os.Create("Pong.hack")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Write each line to the file
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

	// Flush the buffer to ensure everything is written to the file
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
