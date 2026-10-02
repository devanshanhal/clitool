package main

import "fmt"

func main() {
	list := []int{}
	var n int

	fmt.Scan(&n)

	list = append(list, 0)
	list = append(list, 1)

	for i := 2; i < n; i++ {
		list = append(list, list[i-1]+list[i-2])
	}

	fmt.Println(list)
}
