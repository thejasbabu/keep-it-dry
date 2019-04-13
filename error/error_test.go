package error

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsError(t *testing.T) {
	assert.True(t, IsError(errors.New("Error")))
	assert.False(t, IsError(nil))
}
