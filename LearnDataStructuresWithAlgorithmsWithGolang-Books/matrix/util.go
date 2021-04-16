package matrix

var matrix1 = [2][2] int {
	{4,5},
	{1,2},
}

var matrix2 = [2][2] int {
	{6,7},
	{3,4},
}

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var sum [2][2] int
	for i := 0; i < 2; i++ {
		for m := 0; m < 2; m++ {
			sum[i][m] = matrix1[i][m] + matrix2[i][m]
		}
	}
	return sum
}

func subtract(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var difference [2][2]int
	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			difference[l][m] = matrix1[l][m] - matrix2[l][m]
		}
	}
	return difference
}

func multiply(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var product [2][2]int
	for l := 0; l < 2; l++ {
		for m := 0; m < 2; m++ {
			productSum := 0
			for n := 0; n < 2; n++ {
				// 00 00
				// 00 01
				productSum = productSum + matrix1[l][n]*matrix2[n][m]
			}
			product[l][m] = productSum
		}
	}
	return product
}

func transpose(matrix1 [2][2]int) [2][2]int {
	var transMatrix [2][2]int
	for l := 0; l < 2; l++ {
		for m :=0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

func determinant(matrix1 [2][2]int) float32 {
	var det float32
	det = det + ( float32(matrix1[0][0] * matrix1[1][1]) - float32(matrix1[0][1] * matrix1[1][0]) )
	return det
}

func inverse(matrix [2][2]int) [2][2]float32 {
	var det float32
	det = determinant(matrix)
	var invmatrix[2][2] float32

	invmatrix[0][0] = float32(matrix[1][1]) / det
	invmatrix[0][1] = -1 * float32(matrix[0][1]) / det
	invmatrix[1][0] = -1 * float32(matrix[1][0]) / det
	invmatrix[1][1] = float32(matrix[0][0]) / det
	return invmatrix
}
