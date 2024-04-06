package main

import (
	"bufio"
	"math/rand"
	"os"
	"slices"
)

type QuoteGenerator struct {
	quotes              []string
	minTimeBeforeRepeat int // minimum number of times for a quote to be repeated
	lastSeen            []int
}

const MIN_TIME_BEFORE_REPEAT = 10

func NewQuotesGenerator() (*QuoteGenerator, error) {
	quotes, err := readFile("quotes.txt")
	if err != nil {
		return nil, err
	}
	q := QuoteGenerator{
		quotes:              quotes,
		minTimeBeforeRepeat: MIN_TIME_BEFORE_REPEAT,
	}
	// safeguards in case the default value got messed up
	if q.minTimeBeforeRepeat < 0 {
		q.minTimeBeforeRepeat = 0
	}
	if len(quotes) < MIN_TIME_BEFORE_REPEAT {
		q.minTimeBeforeRepeat = len(quotes) - 1
	}
	return &q, nil
}

func readFile(filename string) ([]string, error) {
	quotes := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return quotes, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		q := scanner.Text()
		if q == "" {
			break
		}
		quotes = append(quotes, q)
	}
	return quotes, nil
}

func (q *QuoteGenerator) GetNextQuote() string {
	if len(q.quotes) == 0 {
		return ""
	}
	index := rand.Intn(len(q.quotes))
	for slices.Index(q.lastSeen, index) != -1 {
		index = rand.Intn(len(q.quotes))
	}
	quote := q.quotes[index]
	q.lastSeen = append(q.lastSeen, index)
	if len(q.lastSeen) > q.minTimeBeforeRepeat {
		q.lastSeen = q.lastSeen[1:]
	}
	return quote
}
