package main

import (
	"fmt"
	"main/DubList"
)

func main() {
	fromIntSlice := DubList.New([]int{1, 4, 2, 5, 3})
	fromIntSlice.Append(9)
	fmt.Println("From int slice reversed:   ", fromIntSlice.Reverse())
	fmt.Println("From int slice:            ", fromIntSlice)

	fromStringSlice := DubList.New([]string{"1", "4", "2", "5", "3"})
	fromStringSlice.Append(9)
	fmt.Println("From string slice reversed:", fromStringSlice.Reverse())
	fmt.Println("From string slice:         ", fromStringSlice)
}
