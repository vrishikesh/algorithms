package main

type UFer interface {
	Union(int, int)
	Connected(int, int) bool
	Find(int) int
	Count() int
	Print()
}
