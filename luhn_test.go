package luhn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	b := Check("8703008427912866")
	assert.True(t, b)
}
