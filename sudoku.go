package main

import (
	"fmt"
	"os"
)

func isDuplicate(arr []rune) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == '.' {
			continue
		}

		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}

	return false
}

func isValid(digit rune, rowVal int, colVal int, boardRows [][]rune) bool {
	rowStr := boardRows[rowVal]
	for r := 0; r < 9; r++ {
		if rowStr[r] == digit {
			return false
		}
	}
	for c := 0; c < 9; c++ {
		if boardRows[c][colVal] == digit {
			return false
		}
	}
	starRow := (rowVal / 3) * 3
	starColumn := (colVal / 3) * 3
	for r := starRow; r < starRow+3; r++ {
		for c := starColumn; c < starColumn+3; c++ {
			if boardRows[r][c] == digit {
				return false
			}
		}
	}
	return true
}

func solve(boardRows [][]rune) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if boardRows[r][c] == '.' {
				for i := '1'; i <= '9'; i++ {
					if isValid(i, r, c, boardRows) {
						boardRows[r][c] = i

						if solve(boardRows) {
							return true
						}

						boardRows[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	boardRows := os.Args[1:]

	if len(boardRows) != 9 {
		fmt.Println("Error")
		return
	}

	for _, row := range boardRows {
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}
		runeRow := []rune(row)
		for _, digit := range runeRow {
			if digit != '.' && !(digit >= '1' && digit <= '9') {
				fmt.Println("Error")
				return
			}
		}
		if isDuplicate(runeRow) == true {
			fmt.Println("Error")
			return
		}
	}

	board := make([][]rune, 9)
	for idx, row := range boardRows {
		board[idx] = []rune(row)
	}

	if solve(board) {
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				fmt.Printf("%c", board[r][c])
				if c < 8 {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Error")
	}

}
