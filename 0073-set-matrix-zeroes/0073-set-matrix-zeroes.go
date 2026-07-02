package main

func setZeroes(matrix [][]int) {
    col0 := 1
    n := len(matrix)
    if n == 0 {
        return
    }
    m := len(matrix[0])

    // Step 1: Traverse the matrix and mark 0s in the first row and column
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                if j != 0 {
                    matrix[0][j] = 0
                } else {
                    col0 = 0
                }
            }
        }
    }

    // Step 2: Update the matrix cells using the marked row and column
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][j] != 0 {
                if matrix[0][j] == 0 || matrix[i][0] == 0 {
                    matrix[i][j] = 0
                }
            }
        }
    }

    // Step 3: Handle the first row
    if matrix[0][0] == 0 {
        for j := 0; j < m; j++ {
            matrix[0][j] = 0
        }
    }

    // Step 4: Handle the first column
    if col0 == 0 {
        for i := 0; i < n; i++ {
            matrix[i][0] = 0
        }
    }
}
