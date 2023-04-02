package main

import "fmt"

func LagrangeInterpolation(x, y []float64, t float64) float64 {
	n := len(x)
	sum := 0.0
	for i := 0; i < n; i++ {
		prod := y[i]
		for j := 0; j < n; j++ {
			if i != j {
				prod *= (t - x[j]) / (x[i] - x[j])
			}
		}
		sum += prod
	}
	return sum
}

func main() {
	x := []float64{0, 1, 2, 3, 4}
	y := []float64{1, 0, -1, 0, 1}
	t := 2.5
	interpolatedValue := LagrangeInterpolation(x, y, t)
	fmt.Printf("f(%v) = %v", t, interpolatedValue)
}
