package wc_279

import (
	"fmt"
	"testing"
)

func TestBitSet(t *testing.T) {
	obj := Constructor(75)
	obj.Fix(1)
	obj.Fix(70)
	obj.Unfix(2)
	obj.Flip()
	param_4 := obj.All()
	fmt.Printf("param4: %v\n", param_4)
	param_5 := obj.One()
	fmt.Printf("param5: %v\n", param_5)
	param_6 := obj.Count()
	fmt.Printf("param6: %v\n", param_6)
	param_7 := obj.ToString()
	fmt.Printf("param7: %v\n", param_7)
}

func TestPrintinbinary(t *testing.T) {
	n := 74
	fmt.Printf("%0*b\n", 10, n)
}
