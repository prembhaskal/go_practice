package binsearch

import (
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	testMethod(t, 40, 1000000000)
	testMethod(t, 4001232, 1000000000)
	testMethod(t, 1000000000, 1000000000)
	testMethod(t, 1, 1000000000)
}

func testMethod(t *testing.T, bad, total int) {
	isBad := func(num int) bool {
		return num >= bad
	}

	firstbad := firstBadVersion(total, isBad)
	if firstbad != bad {
		t.Errorf("need: %d, got: %d", bad, firstbad)
	}
}
