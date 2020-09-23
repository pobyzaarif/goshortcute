package goshortcute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumber(t *testing.T) {
	// create 100 random number with length = 10
	count := 100
	test := make([]string, count)
	for i := 0; i < count; i++ {
		test[i] = RandomNumber(10)
	}

	maptest := make(map[string]int)
	for k, tes := range test {
		maptest[tes] = k
	}

	assert.Equal(t, count, len(maptest))
}
