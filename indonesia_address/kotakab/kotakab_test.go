package goshortcute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIDKotakab(t *testing.T) {
	assert.Equal(t, KOTAKAB["1101"], "KAB. ACEH SELATAN")
	assert.Equal(t, GetIDKotakab("KAB. ACEH SELATAN"), "1101")
}
