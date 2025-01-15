package main

import "fmt"

func FindRecursionFibonacci(n int) int {
	// рассматриваемый вариант последовательности
	// для первых чисел 0 1
	return Fibo(0, 1, n)
}

func Fibo(pr, cr, n int) int {
	// если считать от 0, то n == 0
	// если считать от 1, то n == 1
	if n == 1 {
		return pr
	}
	pr, cr = cr, pr+cr
	return Fibo(pr, cr, n-1)
}

func main() {
	fmt.Println(FindRecursionFibonacci(8))
}
