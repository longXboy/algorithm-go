package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UnixNano()
	for i := 0; i < 1000; i++ {
		pow(5.0, i)
	}
	fmt.Println(time.Now().UnixNano() - now)
}

func pow(base float64, exp int) (float64, error) {
	if exp == 0 {
		if base == 0 {
			return 0, fmt.Errorf("zero falut")
		} else {
			return 1.0, nil
		}
	} else if exp > 0 {
		return powUnsigned(base, exp), nil
	} else {
		if base == 0 {
			return 0, fmt.Errorf("zero falut")
		}
		exp = -exp
		return 1 / powUnsigned(base, exp), nil
	}
}

func powUnsigned(base float64, exp int) float64 {
	var total float64 = 1.0
	for {
		if exp == 0 {
			break
		}
		if (exp & 1) != 0 {
			total *= base
		}
		base *= base
		exp = exp >> 1
	}
	return total
}
