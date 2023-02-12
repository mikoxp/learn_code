package main

import (
	"fmt"

	stack "com.moles.stack/elements"
)

func main() {
	fmt.Println("Start")
	s := stack.Create(1)
	stack.WriteOut(s)
	s = stack.Add(s, 2)
	stack.WriteOut(s)
	s = stack.Add(s, 3)
	stack.WriteOut(s)
	s = stack.Add(s, 4)
	stack.WriteOut(s)
	s = stack.Take(s)
	stack.WriteOut(s)
	s = stack.Add(s, 5)
	stack.WriteOut(s)

}
