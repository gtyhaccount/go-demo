package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 8, 7, 6, 5, 4, 2, 3, 9}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			a := 0
			if arr[i] > arr[i+j] {
				a = arr[i]
				arr[i] = arr[i+j]
				arr[i+j] = a
			}
		}
	}

	fmt.Println(arr)

}
