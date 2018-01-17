package main

import "fmt"

func main(){

	i := 1
	for i<=3{
		fmt.Println(i)
		i = i +1
	}

	for j:=5; j<7;j++{
		fmt.Println(j)
	}

	for{
		fmt.Println("Hello")
		break
	}

	for n:=10; n<=15; n++{
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}



}
