package shortcute

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSliceContains(t *testing.T) {
	// Input Slice.
	var slicetest []string
	slicetest = append(slicetest, "stringone", "stringtwo")

	// Lists of normal case scenarios.
	listTest := map[string]bool{
		// Case|Input String : Output
		"Tes1|stringone":   true,
		"Tes2|stringthree": false,
	}

	var output bool
	for input, expect := range listTest {
		in := strings.Split(input, "|")
		output = StringSliceContains(slicetest, in[1])

		assert.Equal(t, output, expect, fmt.Sprintf("%s found on %s", in[1], slicetest))
	}
}
