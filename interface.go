package main

import "fmt"

type Animal interface {
    Speak() string
}
type cow struct{

}

func (c cow) Speak() string{
	return "muuuuuuuuu"
}

type dog struct{

}

func (d dog) Speak() string{
	return "guaguaguaguagua"
}

type donky struct{

}

func (k donky) Speak() string{
	return "ihaihaihaihaiha"
}



func main() {
    animals := []Animal{cow{}, dog{}, donky{}}
    for _, animal := range animals {
        sound := animal.Speak()
	fmt.Println(sound)
    }
}
