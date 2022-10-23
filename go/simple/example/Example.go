package example

import "fmt"

func Write() {
	loop()
	array()
}
func loop() {
	fmt.Println("Loop")
	fmt.Println("Simple")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("infinite")
	for {
		fmt.Println("use break")
		break
	}
}
func array() {
	fmt.Println("Array")
	fmt.Println("Init")
	var tab [5]int
	fmt.Println("tab:", tab)
	tab[3] = 5
	fmt.Println("Set tab[3]=5, tab=", tab)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b", b)
	var s []int
	printArray(s)
	s = append(s, 2, 3, 4)
	s = append(s, 2, 3, 4)
	printArray(s)
	makeArray := make([]string, 3)
	fmt.Println(makeArray)
	c := make([]int, len(s))
	copy(c, s)
	fmt.Println(s, c)
	twoD := make([][]int, 3)
	fmt.Println(twoD)

}
func printArray(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
