package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Matrix struct {
	Rows    int
	Columns int
	Mat     [][]int
}

func (matrix Matrix) getRows() int {
	return matrix.Rows
}

func (matrix Matrix) getColumns() int {
	return matrix.Columns
}

func (matrix *Matrix) setElement( i,  j, element int)  {
	matrix.Mat[i][j] = element
}

func (matrix *Matrix) addMatrix( second Matrix)  {
	for i := 0; i < matrix.Rows; i++ {
		for j := 0; j < matrix.Columns; j++ {
			matrix.Mat[i][j] += second.Mat[i][j]
		}
	}
}

func (matrix Matrix)  printAsJson()  {
	// Convert structs to JSON.
	data, err := json.MarshalIndent(matrix,""," ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func main(){
	first:= Matrix{
		Rows:    3,
		Columns: 3,
		Mat:[][]int{
			{2,3,4},
			{4,1,5},
			{8,9,7},
		},
	}
	second:= Matrix{
		Rows:    3,
		Columns: 3,
		Mat:[][]int{
			{4,1,5},
			{8,9,7},
			{2,3,4},
		},
	}

	fmt.Println("The first matrix is :", first.Mat)
	fmt.Println("The number of rows in first matrix :", first.getRows())
	fmt.Println("The number of columns in first matrix :", first.getColumns())

	first.setElement(2,1,5)

	first.printAsJson()
	fmt.Println("The second matrix is :", second.Mat)

	first.addMatrix(second)

	first.printAsJson()
}

