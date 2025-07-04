package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	sayHello()
	sayHello2("John")
	fmt.Println(sayHello3("John"))

	fmt.Println(randomWithRange(2, 10))

	area, circumference := calculate(7)
	fmt.Println(area)
	fmt.Println(circumference)

	area2, circumference2 := calculate2(7)
	fmt.Println(area2)
	fmt.Println(circumference2)

	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50))
}

func sayHello() {
	fmt.Println("Hello world")
}	

// fungsi dengan parameter dalam golang
func sayHello2(name string) {
	fmt.Println("Hello", name)
}


// fungsi dengan parameter dan return value dalam golang
func sayHello3(name string) string {
	return "Hello" + name
}

// Penggunaan Fungsi rand.New() dalam golang
func randomWithRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Fungsi dengan Multiple Return Value dalam golang
func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d / 2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

// Fungsi Dengan Predefined Return Value dalam golang
func calculate2(d float64) (area, circumference float64) {
	area = math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi * d

	return
}

// Fungsi dengan Variadic Parameter dalam golang
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}