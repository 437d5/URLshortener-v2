package shortener

import (
	"math/rand"
	"time"
)

const Alphabet = "qwertyuiopasdfghjkzxcvbnmQWERTYUPASDFGHJKLZXCVBNM123456789"
const AlphabetLen = len(Alphabet)
const TokenLen = 6

func CreateToken(alphabet string, tokenLen int, alphabetLen int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	token := make([]byte, tokenLen)
	for i := range token {
		token[i] = alphabet[rand.Intn(alphabetLen)]
	}
	return string(token)
}
