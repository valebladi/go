package main
import "fmt"

func sumArray(arr[]int) int {
    var sum int
    for i:=0; i<len(arr); i++{
        sum = sum + arr[i]
    }
    return sum
}

func main() {
    var n, m int

    fmt.Scan(&n)
    array := make([]int, n)

    for j:=0; j<n; j++{
        fmt.Scan(&m)
        array[j] = m
    }

    suma := sumArray(array)

    fmt.Println(suma)
}
