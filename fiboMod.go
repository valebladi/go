package main
import "fmt"

func fibonacci(n uint64, m uint64) (uint64, uint64) {
    o := n + (m*m)
    return m,o
}

func main(){
    var n, m uint64
    var o int
    fmt.Scan(&n, &m, &o)
    for j:=1; j<o; j++ {
        n,m = fibonacci(n,m)
    }
    fmt.Println(n)

}
