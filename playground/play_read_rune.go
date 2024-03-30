package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input string
	var cursorPos int

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		switch char {
		case '\b': // Backspace
			if cursorPos > 0 {
				input = input[:cursorPos-1] + input[cursorPos:]
				cursorPos--
			}
		case '\x7f': // Delete
			if cursorPos < len(input) {
				input = input[:cursorPos] + input[cursorPos+1:]
			}
		case '\x1b': // Start of an escape sequence
			next, _ := reader.Peek(2)
			if len(next) == 2 && next[0] == '[' {
				switch next[1] {
				case 'A': // Up arrow
					fmt.Println("Up arrow pressed")
				case 'B': // Down arrow
					fmt.Println("Down arrow pressed")
				case 'C': // Right arrow
					cursorPos = min(cursorPos+1, len(input))
				case 'D': // Left arrow
					cursorPos = max(cursorPos-1, 0)
				}
				reader.Discard(2) // Discard the escape sequence
			}
		case '\r', '\n': // Enter
			fmt.Println("Enter pressed:", input)
			input = ""
			cursorPos = 0
		default:
			input = input[:cursorPos] + string(char) + input[cursorPos:]
			cursorPos++
		}

		fmt.Printf("\r%s", input) // Print the input string, overwriting the current line
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
