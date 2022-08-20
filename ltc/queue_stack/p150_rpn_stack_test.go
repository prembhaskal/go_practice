package queue_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalRPN(t *testing.T) {
	var tokens []string
	tokens = []string{"2", "3", "+"}
	act := evalRPN(tokens)
	assert.Equal(t, 5, act)

	tokens = []string{"4", "13", "5", "/", "+"}
	act = evalRPN(tokens)
	assert.Equal(t, 6, act)
}
