package main
import "fmt"

func sumArray(arr [5]int) int {
    var sum int
    for i:=0; i<len(arr); i++{
        sum = sum + arr[i]
    }
    arr[len(arr)-1] = 15
    return sum
}

func main() {
    // var n, m int

    // fmt.Scan(&n)
    ar := [5]int{1,2,3,4,5}
    // array := make([]int, n)
    //
    // for j:=0; j<n; j++{
    //     fmt.Scan(&m)
    //     array[j] = m
    // }

    sum := sumArray(ar)

    fmt.Println(ar)

    // suma := sumArray(array)
    //
    fmt.Println(sum)
    // fmt.Println(array)
  }
