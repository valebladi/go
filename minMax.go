package main
import ("fmt"
        "math")

func check(arr[] int) (int,int) {
    max := math.MinInt64
    min := math.MaxInt64
    for i:=0; i<5; i++{
        var sum int
        for j:= 0; j<5; j++{
            if j == i{
                continue
            } else{
                sum = sum + arr[j]
            }
        }
        if max<sum{
            max = sum
        }
        if min>sum{
            min = sum
        }
    }

    return min, max
}

func main(){
    var m int
    array := make([]int, 5)
    for j:=0; j<5; j++{
         fmt.Scan(&m)
         array[j] = m
    }
    //fmt.Println(array)
    x,y := check(array)

    fmt.Println(x,y)

}
