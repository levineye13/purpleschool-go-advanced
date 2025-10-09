package main

import (
	"fmt"
)

func sum(nums []int, ch chan int) {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	ch <- sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3

	sumCh := make(chan int, numGoroutines)
	total := 0

	for index := 0; index < numGoroutines; index++ {
		diff := len(arr) / numGoroutines
		startIndex := diff * index
		endIndex := startIndex + diff - 1

		go func() {
			sum(arr[startIndex:endIndex+1], sumCh)
		}()

		total += <-sumCh
	}

	fmt.Println(total)
}
