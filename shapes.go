package main

import (
    "fmt"
  //  "math"
  )

type geometry interface{
  area()  float64
  perimeter() float64
  }

type rectangle struct{
  width, height float64
}

func (r rectangle) area() float64{
  return r.width * r.height
}

func (r rectangle) perimeter() float64{
  return 2*r.width + 2*r.height
}

func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perimeter())
}


func main(){

  rec := rectangle{2,3}
  //fmt.Println(rec.area())
  measure(rec)

}
