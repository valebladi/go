package main

import (
        "math/rand"
        "time"
	      "fmt"
)

func main(){

	variable := diceRoll(7)
	fmt.Println(variable)

}

func other(){
	fmt.Println("Funcion nueva")
}

func otherother{
  fmt.Println("esta es la branch")
}

func diceRoll(size  int) int  {

  fmt.Println(time.Now().UnixNano())
  rand.Seed(time.Now().UnixNano())
  return rand.Intn(size)  +  1
}
