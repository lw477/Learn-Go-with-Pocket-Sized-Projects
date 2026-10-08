package game

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

var errInvalidWordLength = errors.New("invalid guess, word doesn't have the same number of characters as the solution")

func New(input io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(input),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}
	return g
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.Ask()

		if equalRuntSlice(guess, g.solution) {
			fmt.Printf("You won! You found it in %d guess(es)! The word was: %s.\\n", currentAttempt, string(g.solution))
			return
		}

		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb.String())
	}

	fmt.Printf("You've lost! The solution was: %s.\n", string(g.solution))
}

func (g *Game) Ask() []rune {
	for {
		fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.ValidateGuess(guess)

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid: %s\n", err.Error())
			continue
		}
		return guess
	}
}

func (g *Game) ValidateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expectd %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

func equalRuntSlice(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
