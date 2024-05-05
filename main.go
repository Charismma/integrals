package main

import (
	"fmt"
	"math"
)

func formFunction(x float64) float64 {
	y := 1 + x + math.Sin(2*x)
	return y
}

func UseRectangleRule(xmin float64, xmax float64, num_intervals int) float64 {
	dx := (xmax - xmin) / float64(num_intervals)
	var total_area float64 = 0
	x := xmin
	for i := 0; i < num_intervals; i++ {
		total_area = total_area + dx*formFunction(x)
		x = x + dx
	}
	return total_area
}
func UseTrapezoidRule(xmin float64, xmax float64, num_intervals int) float64 {
	dx := (xmax - xmin) / float64(num_intervals)
	var total_area float64 = 0
	x := xmin
	for i := 0; i < num_intervals; i++ {
		total_area = total_area + dx*(formFunction(x)+formFunction(x+dx))/2
		x = x + dx
	}
	return total_area
}
func main() {
	var xmin float64 = 0
	var xmax float64 = 5
	num_intervals := 10
	area := UseRectangleRule(xmin, xmax, num_intervals)
	area_trapez := UseTrapezoidRule(xmin, xmax, num_intervals)
	fmt.Println("Площадь рассчитанная по методу прямоугольников: ", area)
	fmt.Println("Площадь рассчитанная по методу трапеций: ", area_trapez)

}
