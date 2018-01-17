package main

import (
	"fmt"
	"math"
)


func main(){

	var a string = "a"
	var b bool = true
	var c int = 3

	d := "b"
	e := false
	f := 3


	g := c/f
	h := e && b
	i := a + " " + d

	const pi = 3.1416
	sen := math.Sin(90)

	fmt.Println(a,b,c,d,e,f,g,h,i, pi, sen)

}
