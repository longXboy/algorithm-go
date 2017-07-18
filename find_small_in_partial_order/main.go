package main

import (
	"fmt"
)

func main() {

	fmt.Println(find([]int{3, 4, 5, 1, 2}))
	fmt.Println(find([]int{5, 1, 2, 3, 4}))
	fmt.Println(find([]int{4, 5, 1, 2, 3}))
	fmt.Println(find([]int{2, 3, 4, 5, 1}))

	fmt.Println(find([]int{2, 3, 4, 5, 6, 1}))
	fmt.Println(find([]int{3, 4, 5, 6, 1, 2}))
	fmt.Println(find([]int{4, 5, 6, 1, 2, 3}))
	fmt.Println(find([]int{5, 6, 1, 2, 3, 4}))
	fmt.Println(find([]int{6, 1, 2, 3, 4, 5}))

	fmt.Println(find([]int{1, 0, 1, 1, 1}))
	fmt.Println(find([]int{1, 1, 0, 1, 1}))
	fmt.Println(find([]int{1, 1, 1, 0, 1}))
	fmt.Println(find([]int{1, 1, 1, 1, 0}))

}

func find(a []int) int {
	if len(a) == 1 {
		return a[0]
	} else if len(a) == 2 {
		if a[0] > a[1] {
			return a[1]
		} else {
			return a[0]
		}
	}

	middle := len(a) / 2
	if a[middle] < a[middle-1] {
		return a[middle]
	} else if a[middle] > a[middle+1] {
		return a[middle+1]
	} else if a[middle] > a[len(a)-1] {
		return find(a[middle+1:])
	} else {
		return find(a[:middle])
	}
}
