package queue_stack

import "strconv"
import "fmt"

func evalRPN(tokens []string) int {
	stk := newrpnstack()
	for _, tkn := range tokens {
		switch tkn {
		case "+", "-", "*", "/":
			op1 := stk.pop()
			op2 := stk.pop()
			res := 0
			switch tkn {
			case "+":
				res = op2 + op1
			case "-":
				res = op2 - op1
			case "*":
				res = op2 * op1
			case "/":
				res = op2 / op1
			}

			fmt.Printf("token: %s, op1: %d, op2: %d, res: %d\n", tkn, op1, op2, res)

			stk.push(res)
		default:
			intval, err := strconv.Atoi(tkn)
			if err != nil {
				panic(err)
			}

			stk.push(intval)
		}
	}

	return stk.pop()
}

type rpnstack struct {
	ar []int
}

func newrpnstack() *rpnstack {
	return &rpnstack{
		ar: make([]int, 0),
	}
}

func (s *rpnstack) push(val int) {
	s.ar = append(s.ar, val)
}

func (s *rpnstack) pop() int {
	if s.isempty() {
		return -1
	}
	val := s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1]
	return val
}

func (s *rpnstack) isempty() bool {
	return len(s.ar) == 0
}
