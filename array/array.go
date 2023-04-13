package main

import (
	"fmt"
	"math/rand"
	"time"
)

var size int

func main() {
	var q [5][5]int
	gen(q)
	vivod(q)
	str := calc(q)
	fmt.Println("Строка с самым большим числом", str)
	fmt.Print(rand.Intn(5))
}
func gen(q [5][5]int) {

	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		for j := 0; j < 5; j++ {

			q[i][j] = rand.Intn(10)
		}
	}
}
func vivod(q [5][5]int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(q[i][j])
			fmt.Print(" ")
		}
		fmt.Print("\n\n")
	}
}
func calc(q [5][5]int) int {
	var t int
	max := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if q[i][j] > max {
				max = a[i][j]
				t = i
			}
		}
	}
	return k
}
