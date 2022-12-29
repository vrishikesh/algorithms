package union_find

type UFer interface {
	Union(int, int)
	Connected(int, int) bool
	Find(int) int
	Count() int
	Print()
}
