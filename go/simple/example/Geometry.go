package example

import "fmt"

type geometry interface {
	field() int
	perim() int
}

type rect struct {
	a, b int
}

type square struct {
	a int
}

func (r rect) field() int {
	return r.a * r.b
}
func (r rect) perim() int {
	return 2 * (r.a + r.b)
}

func (s square) field() int {
	return s.a * s.a
}
func (s square) perim() int {
	return 4 * s.a
}

func DataGeometry(g geometry) {
	fmt.Println(g)
	fmt.Println(g.field())
	fmt.Println(g.perim())
}
