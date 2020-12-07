package goshortcute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKecamatan(t *testing.T) {
	assert.Equal(t, KECAMATAN["110101"], "BAKONGAN")
	assert.Equal(t, GetIDKecamatan("BAKONGAN"), "110101")
}
