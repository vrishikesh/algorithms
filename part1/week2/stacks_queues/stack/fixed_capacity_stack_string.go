package stack

import "strings"

type FixedCapacityStackOfStrings struct {
	s []string
	n int
}

func NewFixedCapacityStackOfStrings(capacity int) *FixedCapacityStackOfStrings {
	return &FixedCapacityStackOfStrings{
		s: make([]string, capacity),
		n: 0,
	}
}

func (s *FixedCapacityStackOfStrings) IsEmpty() bool {
	return s.n == 0
}

func (s *FixedCapacityStackOfStrings) Push(val string) {
	s.s[s.n] = val
	s.n += 1
}

func (s *FixedCapacityStackOfStrings) Pop() string {
	s.n -= 1
	item := s.s[s.n]
	s.s[s.n] = ""
	return item
}

func (s *FixedCapacityStackOfStrings) String() string {
	var sb strings.Builder

	for i := 0; i < s.n; i++ {
		sb.WriteString(s.s[i] + " ")
	}

	return sb.String()
}
