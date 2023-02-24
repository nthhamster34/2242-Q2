package main

import "fmt"

// Polygon interface defines the method for calculating the volume of a polygon
type Polygon interface {
	Volume() float64
}

// Triangle struct represents a triangle with base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Volume method calculates the volume of the triangle
func (t Triangle) Volume() float64 {
	return (t.Base * t.Height) / 2
}

// RectangularPrism struct represents a rectangular prism with length, width, and height
type RectangularPrism struct {
	Length float64
	Width  float64
	Height float64
}

// Volume method calculates the volume of the rectangular prism
func (r RectangularPrism) Volume() float64 {
	return r.Length * r.Width * r.Height
}

func main() {
	// create a triangle
	t := Triangle{Base: 4, Height: 5}

	// create a rectangular prism
	r := RectangularPrism{Length: 3, Width: 4, Height: 5}

	// calculate the volumes of the polygon using the Polygon interface
	fmt.Println("Triangle volume:", calculateVolume(t))
	fmt.Println("Rectangular prism volume:", calculateVolume(r))
}

// calculateVolume function takes a Polygon interface and returns its volume
func calculateVolume(p Polygon) float64 {
	return p.Volume()
}
