package game

type Board interface {
	Get(row, col int) int
	Add()
	IsOver() bool
	Total() int
	Input(Key)
	Restart()
}

func (b *board) Get(col, row int) int {
	return b.matrix[row][col]
}

type board struct {
	rand   func() uint32
	matrix [][]int
	over   bool
	newRow int
	newCol int
}

func New(seed uint32) Board {
	b := &board{
		matrix: newMatrix(),
		rand:   lcg(1103515245, 12345, 1<<31, seed),
	}

	b.Add()
	b.Add()

	return b
}

func (b *board) Restart() {
	b.matrix = newMatrix()

	b.Add()
	b.Add()
}

func (b *board) Total() int {
	var total int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			total += b.matrix[i][j]
		}
	}

	return total
}

func (b *board) IsOver() bool {
	empty := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	return empty == 0 || b.over
}

// Add : it first finds the empty slots in the board. They are the one with 0 value
// The it places a new cell randomly in one of those empty places
// The new value to put is also calculated randomly
func (b *board) Add() {
	val := 2

	if b.intn(10) > 8 {
		val = 4
	}

	empty := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}

	elementCount := b.intn(empty) + 1

	index := 0

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.matrix[i][j] == 0 {
				index++
				if index == elementCount {
					b.newRow = i
					b.newCol = j
					b.matrix[i][j] = val
					return
				}
			}
		}
	}

	return
}

func (b *board) intn(n int) int {
	if n <= 0 {
		return -1
	}

	return int(b.rand()) % n
}

func lcg(a, c, m, seed uint32) func() uint32 {
	r := seed
	return func() uint32 {
		r = (a*r + c) % m
		return r
	}
}

func newMatrix() [][]int {
	matrix := make([][]int, 0)

	for i := 0; i < 4; i++ {
		matrix = append(matrix, make([]int, 4))
	}

	return matrix
}
