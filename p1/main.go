package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sort(d []int) {
	for i := 0; i < len(d)-1; i++ {
		for j := 0; j < len(d)-i-1; j++ {
			if d[j] > d[j+1] {
				d[j], d[j+1] = d[j+1], d[j]
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("До сортировки:", arr)
	Sort(arr)
	fmt.Println("После сортировки:", arr)
}
