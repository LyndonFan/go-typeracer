# TypeRacer

Implementing Typeracer on the command line.

![Screenshot of the typeracer running as a terminal app](pictures/03_running.png)

## Usage

(The app is built for Mac / Unix for now -- I'll build for others in a bit!)

A screen by screen guide is in [pictures](./pictures) folder.

0. Run `./go-typeracer`
1. Press `'s'` to start.
2. Wait for the countdown to finish.
3. Start typing! You'll see the word to type underlined, and what you've typed below.
4. When finished, admire your time (or be disappointed by it) on the top right. Then press 's' to restart and 'q' to quit.

## UML Diagram

(I'm new to this so might get some things wrong :sweat_smile: don't judge me)

```mermaid
classDiagram
    class Game {
        Sentence
        Word
        Timer
        Start()
        Reset()
    }
    class Timer {
        StartTime time.Time
        EndTime time.Time
        Start()
        End()
        Reset()
        IsRunning() bool
        TimeElapsed() float
    }
    note for WordController "Takes inputs from tcell.EventKey"
    class WordController {
        currentWord []byte
        Cursor int
        CurrentWord() string
        TakeInput(input byte)
        Reset()
    }
    WordController --* Game
    Timer --* Game
    class Display {
        Game
        WordInput
        Display()
    }
    note for Display "prints game state to terminal"
    Game --* Display
    WordController --* Display
```
