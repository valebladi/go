package main
import "fmt"


func fibonacci(i int ) int{
  if i <= 2{
    return 1
  } else {
    return fibonacci(i-1) + fibonacci(i-2)
  }
}



func main(){

  fib := fibonacci(4)

  fmt.Println(fib)
}
