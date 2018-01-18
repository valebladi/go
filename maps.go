package main
import "fmt"


func main(){
	m := make(map[string]int)
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	sl := make([]int, 3)
	sl2 := []int{1,2,3}

	fmt.Println(sl,sl2, m,m2)
	x := m2["b"]
	fmt.Println(x)
	delete(m2, "b")
	fmt.Println(m2)
	m2["b2"] = 2
	fmt.Println(m2)


}
