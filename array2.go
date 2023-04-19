package main

import "fmt"

func main() {
	números := []int{1, 2, 3, 4, 5}
	fmt.Println(números)
	números = append(números[:2], números[3:]...)
	fmt.Println(números)
}
