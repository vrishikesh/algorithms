package stack

import "strings"

type ResizingArrayStackOfStrings struct {
	s []string
	n int
}

func NewResizingArrayStackOfStrings() *ResizingArrayStackOfStrings {
	return &ResizingArrayStackOfStrings{
		s: make([]string, 1),
		n: 0,
	}
}

func (s *ResizingArrayStackOfStrings) IsEmpty() bool {
	return s.n == 0
}

func (s *ResizingArrayStackOfStrings) resize(n int) {
	arr := make([]string, n)
	copy(arr, s.s)
	s.s = arr
}

func (s *ResizingArrayStackOfStrings) Push(val string) {
	if s.n == len(s.s) {
		s.resize(2 * s.n)
	}

	s.s[s.n] = val
	s.n += 1
}

func (s *ResizingArrayStackOfStrings) Pop() string {
	s.n -= 1
	item := s.s[s.n]
	s.s[s.n] = ""

	if s.n > 0 && s.n == len(s.s)/4 {
		s.resize(len(s.s) / 2)
	}

	return item
}

func (s *ResizingArrayStackOfStrings) String() string {
	var sb strings.Builder

	for i := 0; i < s.n; i++ {
		sb.WriteString(s.s[i] + " ")
	}

	return sb.String()
}
