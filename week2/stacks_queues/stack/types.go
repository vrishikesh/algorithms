package stack

type IStack interface {
	IsEmpty() bool
	Push(string)
	Pop() string
}
