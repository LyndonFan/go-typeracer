package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
)

type State int

const (
	STATE_WELCOME State = iota
	STATE_COUNTDOWN
	STATE_RUNNING
	STATE_FINISHED
)

type Game struct {
	quoteGenerator *QuoteGenerator
	sentence       string
	Tokens         []string
	WordController *WordController
	WordIndex      int
	Timer          *Timer
	State          State
}

func NewGame() (*Game, error) {
	q, err := NewQuotesGenerator()
	if err != nil {
		return nil, err
	}
	g := Game{
		quoteGenerator: q,
		WordController: NewWordController(),
		WordIndex:      0,
		Timer:          NewTimer(),
		State:          STATE_WELCOME,
	}
	return &g, nil
}

func (g *Game) Running() bool {
	return g.State == STATE_COUNTDOWN || g.State == STATE_RUNNING
}

func (g *Game) Sentence() string {
	switch g.State {
	case STATE_WELCOME:
		return "Press 's' to start"
	case STATE_FINISHED:
		return g.sentence + "\nPress 's' to restart\nPress 'q' to quit"
	default:
		return g.sentence
	}
}

func (g *Game) RefreshSentence() {
	quote := g.quoteGenerator.GetNextQuote()
	g.sentence = quote
	g.Tokens = strings.Split(quote, " ")
}

const COUNTDOWN_SECONDS int = 3

func (g *Game) Countdown() {
	g.State = STATE_COUNTDOWN
	// okay to override because original sentence is in g.Tokens
	g.sentence = "Countdown: "
	for i := COUNTDOWN_SECONDS; i > 0; i-- {
		g.sentence += fmt.Sprintf("%d...", i)
		time.Sleep(time.Second)
	}
	g.sentence = strings.Join(g.Tokens, " ")
}

func (g *Game) Start() {
	if g.State == STATE_COUNTDOWN || g.State == STATE_RUNNING {
		return
	}
	g.WordIndex = 0
	g.WordController.Reset()
	g.RefreshSentence()
	g.Countdown()
	g.Timer.Reset()
	g.State = STATE_RUNNING
	g.Timer.Start()
}

func (g *Game) Stop() {
	g.Timer.Stop()
	g.State = STATE_FINISHED
	g.WordController.Reset()
}

func (g *Game) TakeInput(input *tcell.EventKey) {
	switch g.State {
	case STATE_WELCOME, STATE_FINISHED:
		if input.Key() == tcell.KeyRune && input.Rune() == 's' {
			go g.Start()
		}
	case STATE_RUNNING:
		g.runGameLoop(input)
	default:
		// do nothing
	}
}

func (g *Game) runGameLoop(input *tcell.EventKey) {
	g.WordController.TakeInput(input)
	word := g.WordController.CurrentWord()
	if !strings.Contains(word, " ") && g.WordIndex < len(g.Tokens)-1 {
		return
	}
	parts := strings.Split(word, " ")
	for len(parts) > 0 && g.WordIndex < len(g.Tokens) && parts[0] == g.Tokens[g.WordIndex] {
		g.WordIndex++
		parts = parts[1:]
		g.WordController.TrimFirstWord()
	}
	if g.WordIndex == len(g.Tokens) {
		g.Stop()
	}
}
