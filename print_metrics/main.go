package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	printMetrics(a, 0)
}

func printMetrics(a [][]int, round int) {
	xStart := round
	yStart := round
	if xStart < len(a)-round {
		for ; xStart < len(a[0])-round; xStart++ {
			fmt.Println(a[yStart][xStart])
		}
		xStart--
	} else {
		return
	}

	if yStart < len(a)-round-1 {
		yStart++
		for ; yStart < len(a)-round; yStart++ {
			fmt.Println(a[yStart][xStart])
		}
		yStart--
	} else {
		return
	}

	if xStart > round {
		xStart--
		for ; xStart > round-1; xStart-- {
			fmt.Println(a[yStart][xStart])
		}
		xStart++
	} else {
		return
	}

	if yStart > round+1 {
		yStart--
		for ; yStart > round-1; yStart-- {
			fmt.Println(a[yStart][xStart])
		}
		yStart++
	} else {
		return
	}
	printMetrics(a, round+1)
}
