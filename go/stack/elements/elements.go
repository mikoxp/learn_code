package elements

import "fmt"

type Element struct {
	Value int
	next  *Element
}

func WriteOut(el Element) {
	fmt.Print(el.Value)
	if el.next != nil {
		fmt.Print(" -> ")
		WriteOut(*el.next)
	} else {
		fmt.Println("")
	}
}
func Create(value int) Element {
	result := Element{Value: value}
	return result
}
func Add(el Element, value int) Element {
	result := Element{Value: value, next: &el}
	return result
}
func Take(el Element) Element {
	return *el.next
}
