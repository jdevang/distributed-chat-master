package utils

import (
	"math/rand"
	"time"
)

func GenerateMinionUrlIdentifier() string {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz"
	const wordLength = 8
	word := make([]byte, wordLength)
	for i := range word {
		word[i] = charset[rand.Intn(len(charset))]
	}

	return string(word)
}
