package main

import (
	"fmt"
	"math"
)

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// The implementation for `square`s.
func (s square) area() float64 {
	return s.sideLength * s.sideLength
}
func (s square) perim() float64 {
	return 4 * s.sideLength
}

// The implementation for `triangle`s.
func (t triangle) area() float64 {
	return .5 * t.base * t.height
}

func (t triangle) perim() float64 {
	return 4 * t.height
}

// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	s := square{sideLength: 5}
	t := triangle{base: 5, height: 10}

	// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.
	measure(r)
	measure(c)
	measure(s)
	measure(t)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}
func (englishBot) getGreeting() string {
	return "Hello!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
