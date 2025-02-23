package main

func createMatrix(rows, cols int) [][]int {
	//creating a 2d slice wohoo, words that will come back to bite me for sure.
	matrix := make([][]int, rows)

	//slice = append(slice, rows)
	for i := 0; i < rows; i++ {
		row := make([]int, cols)
		for j := 0; j < cols; j++ {
			row[j] = i * j
		}
		matrix[i] = row
	}

	return matrix
}

/*
so in some way we support visualization, lets complete the createMatrix func
it takes a number of rows and columns and returns a 2d clice of ints, the value of each cell
is i*j where i and J are the indexes  of the row and column respectively ,
*/
