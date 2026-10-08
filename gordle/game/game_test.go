package game

import (
	"errors"
	"strings"
	"testing"
)

func TestGame_ValidateGuess(t *testing.T) {
	game := New(strings.NewReader(""), "hello", 3)

	tests := []struct {
		name      string
		guess     string
		wantError bool
	}{
		{name: "same length", guess: "WORLD"},
		{name: "too short", guess: "HI", wantError: true},
		{name: "too long", guess: "GORDLE!", wantError: true},
		{name: "unicode characters", guess: "你好呀世界"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := game.ValidateGuess([]rune(tt.guess))

			if tt.wantError {
				if !errors.Is(err, errInvalidWordLength) {
					t.Fatalf("ValidateGuess(%q) error = %v, want %v", tt.guess, err, errInvalidWordLength)
				}
				return
			}

			if err != nil {
				t.Fatalf("ValidateGuess(%q) unexpected error: %v", tt.guess, err)
			}
		})
	}
}

func TestFeedback_String(t *testing.T) {
	type testCase struct {
		name     string
		guess    string
		solution string
		want     feedback
	}

	tests := []testCase{
		{
			name:     "全部猜对，全绿色",
			guess:    "Hello",
			solution: "hello",
			want: feedback{
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
				correctPosition,
			},
		},
		{
			name:     "字符存在，但是位置全部错误",
			guess:    "LEHLO",
			solution: "Hello",
			want: feedback{
				wrongPosition,
				wrongPosition,
				wrongPosition,
				correctPosition,
				correctPosition,
			},
		},
		{
			name:     "所有字母完全不存在，全部灰色",
			guess:    "ABCDE",
			solution: "HELLO",
			want: feedback{
				absentCharacter,
				absentCharacter,
				absentCharacter,
				absentCharacter,
				absentCharacter,
			},
		},
		{
			name:     "重复字母场景 solution:HELLO guess:SMALL",
			guess:    "SMALL",
			solution: "HELLO",
			// solution H E L L O
			// guess    S M A L L
			// S:无 ⬜; M:无 ⬜; A:无 ⬜; 第4位L 匹配索引2的L 🟩；第5位L，solution中L已经被占用 ⬜
			want: feedback{
				absentCharacter,
				absentCharacter,
				absentCharacter,
				correctPosition,
				absentCharacter,
			},
		},
	}

	for _, tc := range tests {
		guessRunes := splitToUppercaseCharacters(tc.guess)
		solRunes := splitToUppercaseCharacters(tc.solution)

		got := computeFeedback(guessRunes, solRunes)

		if !equalFeebback(t, got, tc.want) {
			t.Fatalf("\ncase: %s\nguess: %q, solution: %q\n got: %s\nwant: %s",
				tc.name, tc.guess, tc.solution, got.String(), tc.want.String())
		}

	}
}

func equalFeebback(t *testing.T, a, b feedback) bool {
	t.Helper()
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
