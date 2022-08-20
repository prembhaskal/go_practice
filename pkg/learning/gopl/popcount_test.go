package gopl

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	var x uint64 = 0b1111000011110010111111001111001111110100111110001111001011111010
	s := PopCount(x)
	assert(t, 42, s)
	s = PopCountLoop(x)
	assert(t, 42, s)
	s = PopCountBitShift(x)
	assert(t, 42, s)
	s = PopCountMinusOne(x)
	assert(t, 42, s)
	s = PopCountMathBits(x)
	assert(t, 42, s)
}

func assert(t *testing.T, exp, act int) {
	if exp != act {
		t.Errorf("exp: %d, act: %d", exp, act)
	}
}

func BenchmarkPopCount(b *testing.B) {
	var x uint64 = 0b1111000011110010111111001111001111110100111110001111001011111010
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	var x uint64 = 0b1111000011110010111111001111001111110100111110001111001011111010
	for i := 0; i < b.N; i++ {
		PopCountLoop(x)
	}
}

func BenchmarkPopCountBitshift(b *testing.B) {
	var x uint64 = 0b1111000011110010111111001111001111110100111110001111001011111010
	for i := 0; i < b.N; i++ {
		PopCountBitShift(x)
	}
}

func BenchmarkPopCountMinusOne(b *testing.B) {
	var x uint64 = 0b1111000011110010111111001111001111110100111110001111001011111010
	for i := 0; i < b.N; i++ {
		PopCountMinusOne(x)
	}
}

func BenchmarkPopCountMathsBit(b *testing.B) {
	var x uint64 = 0b1111000011110010111111001111001111110100111110001111001011111010
	for i := 0; i < b.N; i++ {
		PopCountMathBits(x)
	}
}
