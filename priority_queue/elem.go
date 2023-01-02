package main

import (
	"fmt"
)

type Elem[T comparable] struct {
	Score int
	Data  T
}

func (e *Elem[T]) String() string {
	return fmt.Sprintf("Score: %d, Data: %v", e.Score, e.Data)
}
