package main

import (
	"slices"

	"github.com/gdamore/tcell/v2"
)

type WordController struct {
	cursor      int // between characters, just before ith character
	currentWord []rune
}

func NewWordController() *WordController {
	return &WordController{
		cursor:      0,
		currentWord: []rune{},
	}
}

func (wc *WordController) CurrentWord() string {
	return string(wc.currentWord)
}

func (wc *WordController) Cursor() int {
	return wc.cursor
}

func (wc *WordController) Reset() {
	wc.cursor = 0
	wc.currentWord = []rune{}
}

func (wc *WordController) TrimFirstWord() {
	firstSpaceIndex := slices.Index(wc.currentWord, ' ')
	if firstSpaceIndex == -1 {
		return
	}
	wc.currentWord = wc.currentWord[firstSpaceIndex+1:]
	wc.cursor -= firstSpaceIndex + 1
	if wc.cursor < 0 {
		wc.cursor = 0
	}
}

func (wc *WordController) TakeInput(input *tcell.EventKey) {
	mod, key, inputRune := input.Modifiers(), input.Key(), input.Rune()
	isControl := (mod&tcell.ModCtrl != 0) || (mod&tcell.ModMeta != 0)
	switch key {
	case tcell.KeyLeft:
		if isControl {
			wc.cursor = 0
		} else if wc.cursor > 0 {
			wc.cursor--
		}
	case tcell.KeyRight:
		if isControl {
			wc.cursor = len(wc.currentWord)
		} else if wc.cursor < len(wc.currentWord) {
			wc.cursor++
		}
	case tcell.KeyHome:
		wc.cursor = 0
	case tcell.KeyEnd:
		wc.cursor = len(wc.currentWord)
	case tcell.KeyDEL: // just regular backspace
		if wc.cursor == 0 {
			return
		}
		copy(wc.currentWord[wc.cursor-1:len(wc.currentWord)-1], wc.currentWord[wc.cursor:])
		wc.currentWord = wc.currentWord[:len(wc.currentWord)-1]
		wc.cursor--
		return
	case tcell.KeyBackspace, tcell.KeyCtrlU: // ctrl + backspace, cmd + backspace respectively
		if wc.cursor == 0 {
			return
		}
		wc.currentWord = wc.currentWord[:0]
		wc.cursor = 0
		return
	case tcell.KeyDelete:
		if wc.cursor == len(wc.currentWord) {
			return
		}
		if isControl {
			wc.currentWord = wc.currentWord[:wc.cursor]
		} else {
			wc.currentWord = slices.Delete(wc.currentWord, wc.cursor, wc.cursor+1)
		}
	}
	if inputRune == 0 {
		return
	}
	wc.currentWord = slices.Insert(wc.currentWord, wc.cursor, inputRune)
	wc.cursor++
}
