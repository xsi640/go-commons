package strings

import (
	"math/rand"
	"time"
	"unicode"
)

var (
	Letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Number = []rune("1234567890")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Random(n int, letter, number bool) string {
	result := make([]rune, n)
	var arr []rune
	if letter {
		arr = Letter[:]
	}
	if number {
		arr = append(arr, Number...)
	}

	for i := range result {
		if len(arr) > 0 {
			result[i] = arr[rand.Intn(len(arr))]
		} else {
			result[i] = unicode.MaxRune
		}
	}
	return string(result)
}
