package greedy

import (
	"slices"
)

// Approach - Greedy Approach
//
// Try to find minimum idle slots needed to fullfil all criteria
// use max frequency to get the initial idle slots, fill the idle slots with next frequent character
// if 'n' cycle length is less than unique characters present,
// then we won't need any idle slots, that case we just return task length.
func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		freq[tasks[i]-'A']++
	}

	slices.Sort(freq)

	minIdleSlots := (freq[25] - 1) * n // -1 because we don't need idle slots after last occurrence

	// reduce idle slots needed for next characters
	for i := 24; i >= 0; i-- {
		if freq[i] == 0 {
			break
		}
		if freq[i] == freq[25] {
			minIdleSlots = minIdleSlots - (freq[25] - 1) // we don't need slots after last occurrence
		} else {
			minIdleSlots = minIdleSlots - freq[i]
		}
	}

	if minIdleSlots <= 0 {
		return len(tasks)
	}
	return minIdleSlots + len(tasks)
}
