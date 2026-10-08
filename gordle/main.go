package main

import (
	game "gordle/game"
	"os"
)

const maxAttempts = 3

func main() {
	solution := "hello"
	game := game.New(os.Stdin, solution, maxAttempts)
	game.Play()
}
