package main

import (
	"time"

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
	game, err := NewGame()
	if err != nil {
		panic(err)
	}
	runScreenLoop(game, screen)
}

func runScreenLoop(game *Game, screen tcell.Screen) {
	displayEntireGame(game, screen)
	go func() {
		for {
			updateTimer(game, screen)
			time.Sleep(time.Second / 30)
		}
	}()
	go func() {
		for {
			displayEntireGame(game, screen)
			time.Sleep(time.Second / 2)
		}
	}()
	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			game.TakeInput(event)
			if event.Key() == tcell.KeyCtrlC || event.Key() == tcell.KeyEscape {
				return
			}
			if !game.Running() && event.Key() == tcell.KeyRune && event.Rune() == 'q' {
				return
			}
			updateSentence(game, screen)
			updateCursor(game, screen)
		case *tcell.EventResize:
			screen.Sync()
			displayEntireGame(game, screen)
		}
	}
}
