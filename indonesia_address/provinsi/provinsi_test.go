package goshortcute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvinsi(t *testing.T) {
	assert.Equal(t, PROVINSI["11"], "ACEH")
	assert.Equal(t, GetIDProvinsi("ACEH"), "11")
}
