package main
import "fmt"


type person struct{
	name string
	age int
}

func main(){


	p1 := person{name: "boby", age:15}

	p2 := &p1
	p2.age = 20
	fmt.Println(p1,p2)

}
