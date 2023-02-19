package main

import "fmt"

// makes the board with 0s of nxn size
func createboard(n int) [][]float64 {
	a := make([][]float64, n) // initialize a slice of dy slices
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n) // initialize a slice of dx unit8 in each of dy slices
	}
	return a
}

// progresses the probability of the robot on the close side, aka robot 1
func close(board [][]float64, x int, y int, n int) [][]float64 {
	var currentVal float64 = board[x][y]
	var multiple float64 = 1

	if x < n-1 && y < n-1 {
		multiple = 0.5
	} else {
		multiple = 1
	}

	if x < n-1 {
		board[x+1][y] += (float64(multiple) * currentVal)
	}
	if y < n-1 {
		board[x][y+1] += (float64(multiple) * currentVal)
	}
	return board
}

// progresses the value of the robot on the far side aka robot 2
func far(board [][]float64, x int, y int, n int) [][]float64 {
	var currentVal float64 = board[x][y]
	var multiple float64 = 1

	if x > 0 && y > 0 {
		multiple = 0.5
	} else {
		multiple = 1
	}

	if x > 0 {
		board[x-1][y] += (float64(multiple) * currentVal)
	}
	if y > 0 {
		board[x][y-1] += (float64(multiple) * currentVal)
	}
	return board
}

// prints the board
func printed(board [][]float64, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%f  ", board[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	n := 5
	boardclose := createboard(n)
	boardfar := createboard(n)

	boardclose[0][0] = 1
	boardfar[n-1][n-1] = 1

	//7
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			boardclose = close(boardclose, i, j, n)
			boardfar = far(boardfar, n-i-1, n-j-1, n)
			//if i == n-i-1 && j == n-j-1 {
			//	fmt.Println(boardfar[i][j])
			//}
		}
	}

	printed(boardfar, n)

	// from deduction of where the robots could meet the lowest percentage is 25%
	// appollogies havnt had any sleep so i coudnt really pick up on any assistance
}
