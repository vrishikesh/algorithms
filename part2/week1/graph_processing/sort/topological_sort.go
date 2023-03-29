package sort

import (
	"fmt"
	"graph_processing/types"
	"stacks_queues/generic_stack"
	"strings"
)

type TopologicalSort struct {
	marked []bool
	stack  *generic_stack.LinkedStack[int]
}

func NewTopologicalSort(g types.Graph) *TopologicalSort {
	stack := generic_stack.NewLinkedStack[int]()

	ts := &TopologicalSort{
		marked: make([]bool, g.V()),
		stack:  stack,
	}

	for v := 0; v < g.V(); v++ {
		if !ts.marked[v] {
			ts.dfs(g, v)
		}
	}

	return ts
}

func (ts *TopologicalSort) dfs(g types.Graph, v int) {
	ts.marked[v] = true
	for _, w := range g.Edges(v) {
		if !ts.marked[w] {
			ts.dfs(g, w)
		}
	}
	ts.stack.Push(v)
}

func (ts *TopologicalSort) ReversePost() *generic_stack.LinkedStack[int] {
	return ts.stack
}

func (ts *TopologicalSort) String() string {
	var s strings.Builder
	for !ts.stack.IsEmpty() {
		s.WriteString(fmt.Sprintf("%v ", ts.stack.Pop()))
	}
	return s.String()
}
