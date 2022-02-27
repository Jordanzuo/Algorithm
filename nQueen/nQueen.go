package nQueen

import (
	"fmt"
	"math"
)

type NQueen struct {
	n           int
	count       int
	matrix      [][]int
	validMap    map[int]int
	printResult bool
}

func NewNQueen(n int, printResult bool) *NQueen {
	obj := &NQueen{
		n:           n,
		matrix:      make([][]int, n, n),
		validMap:    make(map[int]int, n),
		printResult: printResult,
	}

	for i := 0; i < n; i++ {
		obj.matrix[i] = make([]int, n)
	}

	return obj
}

func (this *NQueen) GetValidCount() int {
	this.findQueen(0)
	return this.count
}

func (this *NQueen) findQueen(i int) {
	if i >= this.n {
		if this.printResult {
			this.printMatrix()
		}
		this.count++
		return
	}

	for j := 0; j < this.n; j++ {
		if this.isValid(i, j) {
			this.matrix[i][j] = 1
			this.validMap[i] = j
			this.findQueen(i + 1)
			this.matrix[i][j] = 0
			delete(this.validMap, i)
		}
	}
}

func (this *NQueen) isValid(row, col int) bool {
	for i := 0; i < row; i++ {
		j := this.validMap[i]
		if row == i || col == j || row-i == int(math.Abs(float64(col-j))) {
			return false
		}
	}

	return true
}

func (this *NQueen) printMatrix() {
	fmt.Println("Print matrix")
	for i := 0; i < this.n; i++ {
		fmt.Println(this.matrix[i])
	}
	fmt.Println()
}
