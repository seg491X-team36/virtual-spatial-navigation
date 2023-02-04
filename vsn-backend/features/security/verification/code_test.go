package verification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock RandFunc that increments starting at 0
func IncrementRandFunc() RandFunc {
	i := 0
	return func(max int) int {
		temp := i
		i += 1
		return temp
	}
}

func TestCodeGenerator(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		codes := CodeGenerator{
			Characters: DefaultCharacters,
			Length:     6,
			Rand:       IncrementRandFunc(),
		}

		code := codes.Generate()
		assert.Equal(t, "ABCDEF", code)

		code = codes.Generate()
		assert.Equal(t, "GHIJKL", code)
	})

	t.Run("random", func(t *testing.T) {
		codes := CodeGenerator{
			Characters: DefaultCharacters,
			Length:     4,
			Rand:       CryptoRandFunc,
		}

		code := codes.Generate()
		assert.Len(t, code, 4)
		t.Log(code)
	})

}
