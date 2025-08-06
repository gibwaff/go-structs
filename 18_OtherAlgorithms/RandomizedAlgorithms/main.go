package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// monteCarloPi считает число Пи методом Монте-Карло
func monteCarloPi(iterations int) float64 {
	rand.Seed(time.Now().UnixNano())
	insideCircle := 0

	for i := 0; i < iterations; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y <= 1 {
			insideCircle++
		}
	}

	return 4.0 * float64(insideCircle) / float64(iterations)
}

func main() {
	n := 10_000_000 // количество итераций
	pi := monteCarloPi(n)
	fmt.Printf("Оценка числа Пи при %d итерациях: %.6f\n", n, pi)
	fmt.Printf("Погрешность: %.6f\n", math.Abs(math.Pi-pi))
}
