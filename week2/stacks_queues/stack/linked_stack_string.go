package stack

type LinkedStackOfStrings struct {
	first *node
}

func NewLinkedStackOfStrings() *LinkedStackOfStrings {
	return &LinkedStackOfStrings{}
}

func (s *LinkedStackOfStrings) IsEmpty() bool {
	return s.first == nil
}

func (s *LinkedStackOfStrings) Push(val string) {
	second := s.first
	s.first = &node{Val: val, Next: second}
}

func (s *LinkedStackOfStrings) Pop() string {
	val := s.first.Val
	s.first = s.first.Next
	return val
}

func (s *LinkedStackOfStrings) String() string {
	curr := s.first
	var sb string

	for curr != nil {
		sb = curr.Val + " " + sb
		curr = curr.Next
	}

	return sb
}
