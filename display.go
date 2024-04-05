package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

const COMBINING_UNDERLINE rune = 0x0332

func displayEntireGame(game *Game, screen tcell.Screen) {
	screen.Clear()
	width, height := screen.Size()
	displayText(screen, 0, 1, game.Sentence())
	timeElapsed := int(game.Timer.Duration().Seconds())
	displayText(screen, width-5, 0, fmt.Sprintf("%02d:%02d", timeElapsed/60, timeElapsed%60))
	displayWordController(screen, 0, height-2, game.WordController)
	screen.Show()
}

func updateSentence(game *Game, screen tcell.Screen) {
	if !game.Running() {
		return
	}
	displayText(screen, 0, 1, game.Sentence())
	screen.Show()
}

func updateCursor(game *Game, screen tcell.Screen) {
	if !game.Running() {
		return
	}
	width, height := screen.Size()
	// clear out lines
	displayText(screen, 0, height-2, strings.Repeat(" ", width))
	// needed in case flow over to next line
	displayText(screen, 0, height-1, strings.Repeat(" ", width))
	displayWordController(screen, 0, height-2, game.WordController)
	screen.Show()
}

func updateTimer(game *Game, screen tcell.Screen) {
	if !game.Running() {
		return
	}
	width, _ := screen.Size()
	timeElapsed := int(game.Timer.Duration().Seconds())
	displayText(screen, width-5, 0, fmt.Sprintf("%02d:%02d", timeElapsed/60, timeElapsed%60))
	screen.Show()
}

func displayText(s tcell.Screen, startX, startY int, word string) {
	x, y := startX, startY
	width, _ := s.Size()
	for _, r := range word {
		if r == '\n' {
			x = startX
			y++
			continue
		}
		s.SetContent(x, y, r, nil, tcell.StyleDefault)
		x++
		if x >= width {
			x = startX
			y++
		}
	}
}

func displayWordController(s tcell.Screen, startX, startY int, wc *WordController) {
	word := wc.CurrentWord()
	cur := wc.Cursor()
	width, _ := s.Size()
	x, y := startX, startY
	for i, r := range []rune(word) {
		if i == cur {
			s.SetContent(x, y, r, []rune{COMBINING_UNDERLINE}, tcell.StyleDefault)
		} else {
			s.SetContent(x, y, r, nil, tcell.StyleDefault)
		}
		x++
		if x >= width {
			x = startX
			y++
		}
	}
	if cur == len(word) {
		s.SetContent(x, y, '_', nil, tcell.StyleDefault)
	}
}
