package main
import (
    "fmt"
    "math"
)

func sumMatrix(mat[][]int) (int, int) {
    var sum1, sum2 int
    for i:=0; i<len(mat); i++{
        sum1 = sum1 + mat[i][i]
    }

    j := len(mat) -1
    for i:=0; i<len(mat); i++{
        sum2 = sum2 + mat[i][j]
        j = j - 1
    }

    return sum1, sum2
}

func main() {
    var n,m int
    fmt.Scan(&n)

    //var matrix [n][n]int
    matrix := make([][]int, n)


    for i := 0; i < n; i++ {
        matrix[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&m)
            matrix[i][j] = m
        }
    }

    suma1, suma2 := sumMatrix(matrix)
    diag := float64(suma1-suma2)
    total := math.Abs(diag)

    fmt.Println(total)
}
