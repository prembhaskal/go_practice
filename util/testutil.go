package util

import(
	"testing"
)

func AssertInt(t *testing.T, exp, act int) {
	if exp != act {
		t.Errorf("not match, exp: %d, act:%d", exp, act)
	}
}