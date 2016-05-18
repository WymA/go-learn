package main

import (
	"fmt"
	"math"
)

// Vertex struct: 2D vertex
type Vertex struct {
	x, y float64
}

// MyFloat wrap float64
type MyFloat float64

// Main function
func main() {

	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// Abs means absolute method in math
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Abs on Myfloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Scale func that scales x, y value in vertex
func (v *Vertex) Scale(f float64) {
	v.x *= 10
	v.y *= 10
}

// ScaleFunc function
func ScaleFunc(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
