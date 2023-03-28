package types

type Graph interface {
	AddEdge(v, w int)
	V() int
	E() int
	Edges(v int) []int
}
