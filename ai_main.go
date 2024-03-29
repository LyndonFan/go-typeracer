package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[1;34m"
	colorOrange = "\033[1;33m"
)

func main() {
	sentence := "The quick brown fox jumps over the lazy dog"
	fmt.Print("\033[H\033[2J") // clear the terminal
	fmt.Println(sentence)
	fmt.Println("\nStart typing below:")

	reader := bufio.NewReader(os.Stdin)
	typedSentence := ""
	startTime := time.Now()
	wordCount := 0

	for {
		typed, _ := reader.ReadString('\n')
		typed = strings.TrimSpace(typed)
		words := strings.Fields(typed)

		for _, word := range words {
			expectedWord := strings.Fields(sentence)[wordCount]
			if word == expectedWord {
				wordCount++
				typedSentence += word + " "
				fmt.Print("\033[H\033[2J") // clear the terminal
				fmt.Println(sentence)
				fmt.Println("\nStart typing below:")
				fmt.Println(colorBlue + expectedWord + colorReset + " " + colorOrange + typedSentence + colorReset)
			}
		}

		if wordCount == len(strings.Fields(sentence)) {
			break
		}
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("\nTime taken: %02d:%02d\n", int(elapsedTime.Minutes()), int(elapsedTime.Seconds())%60)
}
