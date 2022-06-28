package wc_279

import (
	"fmt"
	"strings"
)

// TODO - complete this.
// HINT - we don't have to really implement bit operations, keep track of all flips, counts etc.
// https://leetcode.com/contest/weekly-contest-279/problems/design-bitset/
type Bitset struct {
	w  []int64
	wl int
}

func Constructor(size int) Bitset {
	n := size / 64
	n = n + 1
	wl := size % 64

	w := make([]int64, n)
	return Bitset{
		w:  w,
		wl: wl,
	}
}

func (this *Bitset) Fix(idx int) {
	// find w idx
	widx := idx / 64
	nidx := idx % 64

	// var fix int64
	fix := int64(1) << nidx
	this.w[widx] = this.w[widx] | fix
}

func (this *Bitset) Unfix(idx int) {
	// find w idx
	widx := idx / 64
	nidx := idx % 64

	var fix int64
	fix = 1 << nidx
	fix = ^fix

	this.w[widx] = this.w[widx] & fix
}

func (this *Bitset) Flip() {
	for i := 0; i < len(this.w); i++ {
		this.w[i] = ^this.w[i]
	}
}

func (this *Bitset) All() bool {
<<<<<<< HEAD
	return false
=======

>>>>>>> eff6115312ba137f186897e55005d208ebf1cc9f
}

func (this *Bitset) One() bool {
	return false
}

func (this *Bitset) Count() int {
	return 0
}

func (this *Bitset) ToString() string {
	var str strings.Builder
	i := 0
	for _, v := range this.w {
		if i == 0 {
			str.WriteString(fmt.Sprintf("%0*b", this.wl, v))
		} else {
			str.WriteString(fmt.Sprintf("%064b", v))
		}
		i++
	}
	return str.String()
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
