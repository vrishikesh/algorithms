package main

type node[T comparable] struct {
	Item T
	Next *node[T]
}
