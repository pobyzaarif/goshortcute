package goshortcute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIDKodepos(t *testing.T) {
	assert.Equal(t, KODEPOS["110101"], []string{"23773"})
	assert.Equal(t, GetIDKodepos("24514"), "110417")
}
