# TypeRacer

Implementing Typeracer on the command line.

## Usage

```bash
./typeracer start
```

Afterwards, you'll be presented with a quote to type and finish.

```bash
Know the enemy and know yourself; in a hundred battles you will never be defeated.


```

## UML Diagram

(I'm new to this so might get some things wrong 😅 don't judge me)

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
    note for WordController "Takes inputs from bufio"
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
