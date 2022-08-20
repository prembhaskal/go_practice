package gopl

import "math/bits"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	s := 0
	for i := 0; i < 8; i++ {
		s += int(pc[byte(x>>(i*8))])
	}
	return s
}

func PopCountBitShift(x uint64) int {
	s := 0
	for x != 0 {
		s += int(x & 1)
		x = x >> 1
	}

	return s
}

func PopCountMinusOne(x uint64) int {
	// x&(x-1) clears rightmost 1.
	s := 0
	for x > 0 {
		x = x & (x - 1)
		s++
	}

	return s
}

func PopCountMathBits(x uint64) int {
	return bits.OnesCount64(x)
}
