package function

func Add(a int, b int) int {
	return a + b
}

func Factor(n int) int {
	if n <= 1 {
		return 1
	}
	return Factor(n-1) * n
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	result := 1
	before := 1
	tmp := 0
	for i := 3; i <= n; i++ {
		tmp = result
		result += before
		before = tmp
	}
	return result
}

func EratosthenesSieve(n int) []int {
	numbersNoFirst := make([]bool, n)
	result := make([]int, 0)
	if n < 1 {
		return result
	}
	numbersNoFirst[0] = true
	for i := 2; i <= n; i++ {
		if numbersNoFirst[i-1] {
			continue
		}
		j := 2
		for i*j <= n {
			numbersNoFirst[(i*j)-1] = true
			j++
		}
	}

	for i := 0; i < n; i++ {
		if !numbersNoFirst[i] {
			result = append(result, i+1)
		}
	}
	return result[:]
}
