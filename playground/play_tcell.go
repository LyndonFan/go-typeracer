package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
	screen.Clear()
	screen.Show()
	x, y := 0, 0
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()
			// show this on the screen
			// "EventKey Modifiers: %d Key: %d Rune: %d\n", mod, key, ch
			content := []rune(fmt.Sprintf("EventKey Modifiers: %d Key: %d Rune: %d\n", mod, key, ch))
			screen.SetContent(x, y, content[0], content[1:], tcell.StyleDefault)
			// if is esc: exit
			if key == tcell.KeyEscape {
				return
			}
		case *tcell.EventResize:
			screen.Sync()
		}
		x++
		y++
		w, h := screen.Size()
		x = x % w
		y = y % h
		screen.Show()
	}
}
