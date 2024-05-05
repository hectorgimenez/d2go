package nip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNIPFile(t *testing.T) {
	path := "mock/rare.nip"

	rules, err := ParseNIPFile(path)
	assert.NoError(t, err)
	assert.NotEmpty(t, rules)
}
