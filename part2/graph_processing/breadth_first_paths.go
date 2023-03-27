package main

import (
	"fmt"
	"stacks_queues/generic_queue"
	"strings"
)

type BreadthFirstPaths struct {
	edgeTo []int
	distTo []int
	s      int
}

func NewBreadthFirstPaths(g *MyGraph, s int) *BreadthFirstPaths {
	b := &BreadthFirstPaths{
		edgeTo: make([]int, g.V()),
		distTo: make([]int, g.V()),
		s:      s,
	}

	for i := 0; i < g.V(); i++ {
		b.distTo[i] = -1
	}

	b.bfs(g, s)

	return b
}

func (b *BreadthFirstPaths) bfs(g *MyGraph, s int) {
	q := generic_queue.NewLinkedQueue[int]()
	b.distTo[s] = 0
	q.Enqueue(s)

	for !q.IsEmpty() {
		v := q.Dequeue()
		edges, ok := g.Edges(v)
		if ok {
			for _, w := range edges {
				if b.distTo[w] == -1 {
					b.edgeTo[w] = v
					b.distTo[w] = b.distTo[v] + 1
					q.Enqueue(w)
				}
			}
		}
	}
}

func (b *BreadthFirstPaths) HasPathTo(v int) bool {
	return b.distTo[v] != -1
}

func (b *BreadthFirstPaths) PathTo(v int) []int {
	if !b.HasPathTo(v) {
		return nil
	}

	path := []int{}
	for x := v; x != b.s; x = b.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, b.s)

	return path
}

func (b *BreadthFirstPaths) String() string {
	var s strings.Builder
	for v := range b.distTo {
		s.WriteString(fmt.Sprintf("%v: ", v))
		path := b.PathTo(v)
		for _, p := range path {
			s.WriteString(fmt.Sprintf("%v ", p))
		}
		s.WriteString("\n")
	}
	return s.String()
}
