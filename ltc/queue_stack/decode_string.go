package queue_stack

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	// stack of strings
	// do until stack is not empty.
	// if char != ']', add to stack
	// if char == ']', remove from stack until [ and number there after.
	// replicate and push back to stack
	// 3[ac2[ab]ef]ac => 3[acababef]ac => acababefacababefacababefac

	digits := make(map[string]bool, 0)
	for i := 0; i < 10; i++ {
		digits[digitToString(i)] = true
	}

	stk := newdecodestack()
	for _, rn := range s {
		if rn != ']' {

			if rn == '[' {
				// convert pushed digits to numbers and put it back
				var tmpstr strings.Builder
				for {
					headstr := stk.pop()
					if digits[headstr] {
						tmpstr.WriteString(headstr)
					} else {
						stk.push(headstr)
						break
					}
				}

				num, _ := strconv.Atoi(reverseString(tmpstr.String()))
				stk.push(strconv.Itoa(num))
			}

			stk.push(string(rn))

		} else {
			headstr := stk.pop()
			var tempstr strings.Builder
			for headstr != "[" {
				tempstr.WriteString(headstr)
				headstr = stk.pop()
			}
			repeat, _ := strconv.Atoi(stk.pop())

			var newarr strings.Builder
			for ; repeat > 0; repeat-- {
				newarr.WriteString(tempstr.String())
			}
			stk.push(newarr.String())
			// a -> bd -> c       ["c", "bd", "a"]
		}
	}

	var finalstr strings.Builder
	for !stk.isempty() {
		finalstr.WriteString(stk.pop())
	}

	return reverseString(finalstr.String())
}

func digitToString(digit int) string {
	return strconv.Itoa(digit)
}

func reverseString(input string) string {
	chars := []rune(input)
	n := len(chars)
	for i := 0; i < n/2; i++ {
		chars[i], chars[n-i-1] = chars[n-i-1], chars[i]
	}

	return string(chars)
}

type decodestack struct {
	ar []string
}

func newdecodestack() *decodestack {
	return &decodestack{
		ar: make([]string, 0),
	}
}

func (s *decodestack) push(str string) {
	s.ar = append(s.ar, str)
}

func (s *decodestack) pop() string {
	if s.isempty() {
		return ""
	}
	val := s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1]
	return val
}

func (s *decodestack) isempty() bool {
	return len(s.ar) == 0
}
