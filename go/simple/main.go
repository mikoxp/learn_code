package main

import (
	"fmt"

	"com.moles/example"
	f "com.moles/functions"
)

func main() {
	var a int = 1
	var b int = 2
	var n int = 5
	var nn int = 20
	fmt.Println("Hello, World!")
	fmt.Printf("%d + %d = %d\n", a, b, f.Add(a, b))
	fmt.Printf("Factor n=%d is %d\n", n, f.Factor(n))
	fmt.Printf("Fibonacci sequence n=%d is %d\n", n, f.Fibonacci(n))
	fmt.Println("--------------------------")
	fmt.Printf("Eratosthenes Sieve n=%d is :\n", nn)
	fmt.Println(f.EratosthenesSieve(nn))
	fmt.Println("--------------------------")
	example.Write()
	example.Figure()

}
