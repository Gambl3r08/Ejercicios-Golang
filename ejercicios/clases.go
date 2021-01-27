package main

import (
	"fmt"
	"math"
)

func main() {
	p := &Punto{3, 4}
	fmt.Println(p.Abs())
}

type Punto struct {
	x, y float64
}

func (p *Punto) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
