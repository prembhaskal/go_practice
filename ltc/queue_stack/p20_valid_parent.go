package queue_stack

func isValid(s string) bool {

	matchMap := make(map[rune]rune, 0)
	matchMap[')'] = '('
	matchMap['}'] = '{'
	matchMap[']'] = '['

	// add char in stack if it belongs if among - (, {, [
	// if char in string is among - ), }, ]
	// pop stack, if it not its pair return false.
	stk := newstk20()
	for _, chr := range s {
		if chr == '(' || chr == '{' || chr == '[' {
			stk.push(chr)
		} else {
			if stk.isempty() {
				return false
			}

			top := stk.pop()
			if top != matchMap[chr] {
				return false
			}

		}
	}

	return stk.isempty()
}

type stck struct {
	ar []rune
}

func newstk20() *stck {
	return &stck{
		ar: make([]rune, 0),
	}
}

func (s *stck) push(val rune) {
	s.ar = append(s.ar, val)
}

func (s *stck) pop() rune {
	if s.isempty() {
		return -1
	}
	val := s.ar[len(s.ar)-1]
	s.ar = s.ar[:len(s.ar)-1]
	return val
}

func (s *stck) isempty() bool {
	return len(s.ar) == 0
}
