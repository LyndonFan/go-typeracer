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
	sentence       string
	tokens         []string
	WordController *WordController
	WordIndex      int
	Timer          *Timer
	State          State
}

func NewGame(sentence string) *Game {
	return &Game{
		sentence:       sentence,
		tokens:         strings.Split(sentence, " "),
		WordController: NewWordController(),
		WordIndex:      0,
		Timer:          NewTimer(),
		State:          STATE_WELCOME,
	}
}

func (g *Game) Running() bool {
	return g.State == STATE_COUNTDOWN || g.State == STATE_RUNNING
}

func (g *Game) Sentence() string {
	switch g.State {
	case STATE_WELCOME:
		return "Press 's' to start"
	case STATE_FINISHED:
		return g.sentence + "\nPress 'r' to restart"
	default:
		return g.sentence
	}
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
	g.sentence = strings.Join(g.tokens, " ")
}

func (g *Game) Start() {
	if g.State != STATE_WELCOME {
		return
	}
	g.WordIndex = 0
	g.WordController.Reset()
	g.Countdown()
	g.Timer.Reset()
	g.State = STATE_RUNNING
	g.Timer.Start()
}

func (g *Game) Stop() {
	g.State = STATE_FINISHED
	g.Timer.Stop()
}

func (g *Game) TakeInput(input *tcell.EventKey) {
	switch g.State {
	case STATE_WELCOME:
		if input.Key() == tcell.KeyRune && input.Rune() == 's' {
			go g.Start()
		}
	case STATE_RUNNING:
		g.runGameLoop(input)
	case STATE_FINISHED:
		if input.Key() == tcell.KeyRune && input.Rune() == 'r' {
			go g.Start()
		}
	default:
		// do nothing
	}
}

func (g *Game) runGameLoop(input *tcell.EventKey) {
	g.WordController.TakeInput(input)
	word := g.WordController.CurrentWord()
	if !strings.Contains(word, " ") && g.WordIndex < len(g.tokens)-1 {
		return
	}
	parts := strings.Split(word, " ")
	for len(parts) > 0 && g.WordIndex < len(g.tokens) && parts[0] == g.tokens[g.WordIndex] {
		g.WordIndex++
		parts = parts[1:]
		g.WordController.TrimFirstWord()
	}
	if g.WordIndex == len(g.tokens) {
		g.Stop()
	}
}
