package slider_puzzle

import (
	"fmt"
	"math"
	"strings"
)

type Matrix [][]int

func (m Matrix) String() string {
	var sb strings.Builder
	N := len(m)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sb.WriteString(fmt.Sprintf("%d ", m[i][j]))
		}
	}

	return sb.String()
}

func (m Matrix) Clone() Matrix {
	N := len(m)
	clone := make(Matrix, N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			clone[i] = append(clone[i], m[i][j])
		}
	}

	return clone
}

type Tile struct {
	Row int
	Col int
}

func (t Tile) String() string {
	return fmt.Sprintf("(%d, %d)", t.Row, t.Col)
}

// GetSides returns the sides of a tile
func (t *Tile) GetSides() [4]Tile {
	return [4]Tile{
		{t.Row + 1, t.Col},
		{t.Row, t.Col + 1},
		{t.Row - 1, t.Col},
		{t.Row, t.Col - 1},
	}
}

type Board struct {
	Goal         Matrix
	M            Matrix
	NeighborList map[string]struct{}
}

func (b Board) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Goal:         %s", b.Goal) + "\n")
	sb.WriteString(fmt.Sprintf("Current:      %s", b.M) + "\n")
	sb.WriteString(fmt.Sprintf("NeighborList: %s", b.NeighborList) + "\n")

	return sb.String()
}

// create a board from an n-by-n array of tiles,
// where tiles[row][col] = tile at (row, col)
func New(m Matrix) *Board {
	N := len(m)
	goal := make(Matrix, N)

	for i, k := 0, 1; i < N; i++ {
		for j := 0; j < N; j++ {
			goal[i] = append(goal[i], k)
			k += 1
		}
	}
	goal[N-1][N-1] = 0

	return &Board{
		Goal:         goal,
		M:            m,
		NeighborList: map[string]struct{}{},
	}
}

// number of tiles out of place
func (b *Board) Hamming(m, n Matrix) int {
	hammingDistance := 0
	N := len(m)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if m[i][j] != n[i][j] {
				hammingDistance += 1
			}
		}
	}
	return hammingDistance
}

// sum of Manhattan distances between tiles and goal
func (b *Board) Manhattan(m, n Matrix) int {
	var manhattanDistance float64 = 0
	boardMap := map[int][2]int{}
	N := len(m)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			tile := m[i][j]
			boardMap[tile] = [2]int{i, j}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if m[i][j] != n[i][j] {
				tile := n[i][j]
				loc := boardMap[tile]
				x, y := loc[0], loc[1]
				dx, dy := float64(x-i), float64(y-j)
				manhattanDistance += math.Abs(dx) + math.Abs(dy)
			}
		}
	}
	return int(manhattanDistance)
}

// is this board the goal board?
func (b *Board) IsGoal(m Matrix) bool {
	return b.Equals(m, b.Goal)
}

// does this board equal y?
func (b *Board) Equals(m, n Matrix) bool {
	N := len(m)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if m[i][j] != n[i][j] {
				return false
			}
		}
	}

	return true
}

// IsNeighbor returns true if the given matrix is a neighbor of the current board
func (b *Board) IsNeighbor(m Matrix) bool {
	_, ok := b.NeighborList[m.String()]
	return ok
}

// AddToNeighbors adds the given matrix to the neighbor list
func (b *Board) AddToNeighbors(m Matrix) {
	b.NeighborList[m.String()] = struct{}{}
}

// a board that is obtained by exchanging any pair of tiles
func (b *Board) Twin(m Matrix, t1, t2 Tile) {
	if b.IsGoal(m) {
		return
	}

	if t1.Row == t2.Row && t1.Col == t2.Col {
		return
	}

	if b.IsTileInMatrix(m, t1) && b.IsTileInMatrix(m, t2) {
		b.AddToNeighbors(m)
		m[t2.Row][t2.Col], m[t1.Row][t1.Col] = m[t1.Row][t1.Col], m[t2.Row][t2.Col]
	}
}

// IsInMatrix returns true if the tile is in the matrix
func (b *Board) IsTileInMatrix(m Matrix, t Tile) bool {
	return t.Row >= 0 && t.Row < len(m) && t.Col >= 0 && t.Col < len(m)
}

// GetEmptyTile returns the empty tile
func (b *Board) GetEmptyTile() Tile {
	N := len(b.M)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if b.M[i][j] == 0 {
				return Tile{Row: i, Col: j}
			}
		}
	}

	return Tile{}
}

func (b *Board) Start() {
	for !b.IsGoal(b.M) {
		tile := b.GetEmptyTile()
		sides := tile.GetSides()
		fmt.Printf("sides: %v\n", sides)
		distances := make([]int, 0)
		matrices := make([]Matrix, 0)

		for _, side := range sides {
			if b.IsTileInMatrix(b.M, side) {
				clone := b.M.Clone()

				b.Twin(clone, side, tile)
				// fmt.Printf("Clone:   %s\n", clone)

				distance := b.Hamming(clone, b.Goal)
				// fmt.Printf("Hamming: %d\n", distance)

				distances = append(distances, distance)
				matrices = append(matrices, clone)
			}
		}

		i, d := min(distances...)
		fmt.Printf("min distance: %v\n", d)
		// fmt.Printf("min matrix: %v\n", matrices[i])
		// fmt.Printf("NeighborList: %v\n\n", b.NeighborList)
		b.M = matrices[i]
	}

	fmt.Printf("Goal: %s\n", b.M)
}

func min(x ...int) (int, int) {
	min := x[0]
	index := 0
	for i, v := range x {
		if v < min {
			min = v
			index = i
		}
	}
	return index, min
}
