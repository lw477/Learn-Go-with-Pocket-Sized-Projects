package game

import "strings"

// hint 每个字母的状态 iota枚举
type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

type feedback [] hint

func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜"
	case wrongPosition:
		return "🟨"
	case correctPosition:
		return "🟩"
	default:
		return "❌"
	}
}


func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	for i := range guess {
		if guess[i] == solution[i]{
			result[i] = correctPosition
			used[i] = true
		} else {
			result[i] = absentCharacter
		}
	}

	for guessIdx := range guess {
		if result[guessIdx] == correctPosition {
			continue
		}
		for solIdex := range solution {
			if !used[solIdex] && solution[solIdex] == guess[guessIdx] {
				result[guessIdx] = wrongPosition
				used[solIdex] = true
				break
			}
		}
	}
	return result
}