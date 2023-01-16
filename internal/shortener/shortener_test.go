package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomSequence(t *testing.T) {
	randomStr := GetRandomSequence(6)
	assert.NotNil(t, randomStr)
}
