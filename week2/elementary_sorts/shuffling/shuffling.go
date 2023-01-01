package shuffling

import "math/rand"

func Shuffle(slice []int) {
	N := len(slice)

	for i := 0; i < N; i++ {
		r := rand.Intn(N-i) + i
		slice[i], slice[r] = slice[r], slice[i]
	}
}

func Generate(capacity int) []int {
	var slice = make([]int, capacity)

	for i := 0; i < capacity; i++ {
		slice[i] = i
	}

	Shuffle(slice)

	return slice
}
