package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()
	sentence := "The quick brown fox jumps over the lazy dog."
	game := NewGame(sentence)
	game.Start()
	runScreenLoop(game, screen)
}

func runScreenLoop(game *Game, screen tcell.Screen) {
	displayText(screen, 0, 0, game.Sentence)
	displayText(screen, 0, 1, "(start typing)")
	for game.Timer.Running {
		event := screen.PollEvent()
		screen.Clear()
		switch event := event.(type) {
		case *tcell.EventKey:
			game.InputChannel <- event
			if event.Key() == tcell.KeyCtrlC || event.Key() == tcell.KeyEscape {
				return
			}
			displayText(screen, 0, 0, game.Sentence)
			displayWord(screen, 0, 1, game.WordController)
			cursorMessage := fmt.Sprintf("Cursor: %d\n", game.WordController.Cursor())
			displayText(screen, 0, 2, cursorMessage)
			// mod, key, r := event.Modifiers(), event.Key(), event.Rune()
			// keyMessage := fmt.Sprintf("EventKey Modifiers: %d Key: %d Rune: %d\n", mod, key, r)
			// displayText(screen, 0, 2, keyMessage)
			// backSpaceMessage := fmt.Sprintf("Backspace Key: %d\n", tcell.KeyBackspace)
			// displayText(screen, 0, 3, backSpaceMessage)
		case *tcell.EventResize:
			screen.Sync()
		}
		screen.Show()
	}
}

func displayText(s tcell.Screen, x, y int, word string) {
	for i, r := range []rune(word) {
		s.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}

func displayWord(s tcell.Screen, x, y int, wc *WordController) {
	word := wc.CurrentWord()
	for i, r := range []rune(word) {
		s.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
	cur := wc.Cursor()
	if cur == len(word) {
		s.SetContent(x+cur, y, '_', nil, tcell.StyleDefault)
	} else {
		s.SetContent(x+cur, y, []rune(word)[cur], []rune{'\u0332'}, tcell.StyleDefault) // add underline
	}
}
