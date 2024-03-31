package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Sentence       string
	tokens         []string
	InputChannel   chan *tcell.EventKey
	WordController *WordController
	Timer          *Timer
}

func NewGame(sentence string) *Game {
	inputChannel := make(chan *tcell.EventKey)
	return &Game{
		Sentence:       sentence,
		tokens:         strings.Split(sentence, " "),
		InputChannel:   inputChannel,
		WordController: NewWordController(),
		Timer:          NewTimer(),
	}
}

func (g *Game) Start() {
	g.Timer.Start()
	go g.runGameLoop()
}

func (g *Game) Stop() {
	g.Timer.Stop()
}

func (g *Game) runGameLoop() {
	defer close(g.InputChannel)
	wordIndex := 0
	for input := range g.InputChannel {
		g.WordController.TakeInput(input)
		word := g.WordController.CurrentWord()
		if !strings.Contains(word, " ") {
			continue
		}
		parts := strings.Split(word, " ")
		for len(parts) > 0 && wordIndex < len(g.tokens) && parts[0] == g.tokens[wordIndex] {
			wordIndex++
			parts = parts[1:]
			g.WordController.TrimFirstWord()
		}
		if wordIndex == len(g.tokens) {
			g.Stop()
			break
		}
	}
}
