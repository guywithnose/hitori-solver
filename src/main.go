package main

import (
	"bufio"
	"fmt"
	"hitori"
	"os"
)

func main() {
	args := os.Args
	var cellData [][]int
	if len(args) != 2 {
		reader := bufio.NewReader(os.Stdin)
		line, error := reader.ReadString('\n')
		for error == nil {
			var row []int
			for _, char := range line {
				if char >= 48 && char <= 57 {
					row = append(row, int(char)-48)
				}
			}

			if len(row) == 9 {
				cellData = append(cellData, row)
			}

			line, error = reader.ReadString('\n')
		}
	} else {
		reader, error := os.Open(args[1])
		if error != nil {
			fmt.Println("Unable to open file:", args[1])
			return
		}

		rowBuffer := make([]byte, 10)
		_, error = reader.Read(rowBuffer)
		for error == nil {
			var row []int
			for _, char := range rowBuffer {
				if char >= 48 && char <= 57 {
					row = append(row, int(char)-48)
				}
			}

			if len(row) != 9 {
				panic("Invalid file format")
			}

			cellData = append(cellData, row)

			_, error = reader.Read(rowBuffer)
		}
	}

	if len(cellData) != 9 {
		panic("Invalid format")
	}

	board := hitori.PopulateBoard(cellData)
	board.Solve()
	fmt.Println(board)
}
