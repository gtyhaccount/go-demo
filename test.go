package main

import (
	"fmt"
	"time"
)

func main() {
	format := "010602150405"

	dateStr := "071930160000"

	t, _ := time.ParseInLocation(format, dateStr, time.Local)
	fmt.Println(t)
}
