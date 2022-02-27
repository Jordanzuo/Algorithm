package nDiagonal

import (
	"encoding/json"
	"fmt"
)

type Diagonal struct {
	n           int
	count       int
	threshold   int
	matrix      [][]int
	validMap    map[string]struct{}
	printResult bool
}

func NewDiagonal(n, threshold int, printResult bool) *Diagonal {
	obj := &Diagonal{
		n:           n,
		threshold:   threshold,
		matrix:      make([][]int, n),
		validMap:    make(map[string]struct{}, 0),
		printResult: printResult,
	}

	for i := 0; i < n; i++ {
		obj.matrix[i] = make([]int, n)
	}

	return obj
}

func (this *Diagonal) GetValidCount() int {
	this.findMatrix(0, 0)
	return this.count
}

func (this *Diagonal) findMatrix(i, j int) {
	if i >= this.n {
		if this.isMatrixValid() {
			fingerPrint := this.fingerPrint()
			if _, exists := this.validMap[fingerPrint]; !exists {
				this.count++
				if this.printResult {
					this.printMatrix()
				}
				this.validMap[fingerPrint] = struct{}{}
			}
		}
		return
	}

	for val := 1; val <= 2; val++ {
		if this.isValid(i, j, val) {
			this.matrix[i][j] = val
		}
		if j != this.n-1 {
			this.findMatrix(i, j+1)
		} else {
			this.findMatrix(i+1, 0)
		}
		this.matrix[i][j] = 0
	}
}

func (this *Diagonal) printMatrix() {
	fmt.Println("Print matrix")
	for i := 0; i < this.n; i++ {
		fmt.Println(this.matrix[i])
	}
	fmt.Println()
}

func (this *Diagonal) isMatrixValid() bool {
	num := 0
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.n; j++ {
			if this.matrix[i][j] > 0 {
				num++
			}
		}
	}

	return num >= this.threshold
}

func (this *Diagonal) isValid(row, col, val int) bool {
	// fmt.Printf("row: %d, col: %d, val: %d\n", row, col, val)
	if val == 1 {
		if row-1 >= 0 && this.matrix[row-1][col] == 2 {
			return false
		}
		if row-1 >= 0 && col+1 <= this.n-1 && this.matrix[row-1][col+1] == 1 {
			return false
		}
		if col-1 >= 0 && this.matrix[row][col-1] == 2 {
			return false
		}
	} else if val == 2 {
		if row-1 >= 0 && col-1 >= 0 && this.matrix[row-1][col-1] == 2 {
			return false
		}
		if row-1 >= 0 && this.matrix[row-1][col] == 1 {
			return false
		}
		if col-1 >= 0 && this.matrix[row][col-1] == 1 {
			return false
		}
	}

	return true
}

func (this *Diagonal) fingerPrint() string {
	bytes, _ := json.Marshal(this.matrix)
	return string(bytes)
}
