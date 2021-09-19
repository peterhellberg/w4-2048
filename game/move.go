package game

func (b *board) left() {
	for i := 0; i < 4; i++ {
		old := b.matrix[i]
		b.matrix[i] = movedRow(old)
	}
}

func (b *board) up() {
	b.reverseRows()
	b.down()
	b.reverseRows()
}

func (b *board) down() {
	b.transpose()
	b.left()
	b.transpose()
	b.transpose()
	b.transpose()
}

func (b *board) right() {
	b.reverse()
	b.left()
	b.reverse()
}

// movedRow simply finds empty elements and filled elements
// it places the filled element in the beginning of the row
// [2 0 3 0] will become [2 3 0 0]
// an empty cell is displayed with 0 value
func movedRow(elems []int) []int {
	nonEmpty := make([]int, 0)

	for i := 0; i < 4; i++ {
		if elems[i] != 0 {
			nonEmpty = append(nonEmpty, elems[i])
		}
	}

	remaining := 4 - len(nonEmpty)

	for i := 0; i < remaining; i++ {
		nonEmpty = append(nonEmpty, 0)
	}

	return mergeElements(nonEmpty)
}

// reverse simply reverses each row of the board
func (b *board) reverse() {
	for i := 0; i < 4; i++ {
		b.matrix[i] = reverseRow(b.matrix[i])
	}
}

// transpose rotates a list
// row becomes cols
// [ 1 2 ]
// [ 3 4 ] becomes
//
// [ 3 1 ]
// [ 4 2 ]
func (b *board) transpose() {
	m := make([][]int, 0)

	for i := 0; i < 4; i++ {
		m = append(m, make([]int, 4))
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = b.matrix[4-j-1][i]
		}
	}

	b.matrix = m
}

// reverseRows reverses the order of lists
// [1 2]
// [3 4] becomes
//
// [3 4]
// [1 2]
func (b *board) reverseRows() {
	m := make([][]int, 0)

	for i := 0; i < 4; i++ {
		m = append(m, make([]int, 4))
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[4-i-1][j] = b.matrix[i][j]
		}
	}

	b.matrix = m
}

// reverseRow reverses a row
func reverseRow(arr []int) []int {
	row := make([]int, 0)

	for i := len(arr) - 1; i >= 0; i-- {
		row = append(row, arr[i])
	}

	return row
}

// mergeElements when a row is moved to left, it merges the element which can
func mergeElements(old []int) []int {
	row := make([]int, len(old))
	row[0] = old[0]

	index := 0

	for i := 1; i < len(old); i++ {
		if old[i] == row[index] {
			row[index] += old[i]
		} else {
			index++
			row[index] = old[i]
		}
	}

	return row
}
