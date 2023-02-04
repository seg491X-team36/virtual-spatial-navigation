package verification

import (
	"crypto/rand"
	"math/big"
)

var DefaultCharacters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type CodeGenerator struct {
	Characters []rune   // characters included in the codes
	Length     int      // number of characters in each code
	Rand       RandFunc // random number generator
}

// generate a new code
func (c *CodeGenerator) Generate() string {
	code := make([]rune, c.Length)
	for i := range code {
		next := c.Rand(len(c.Characters))
		code[i] = c.Characters[next]
	}
	return string(code)
}

type RandFunc func(max int) int

func CryptoRandFunc(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}
