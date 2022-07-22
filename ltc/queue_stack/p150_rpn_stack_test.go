package queue_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
