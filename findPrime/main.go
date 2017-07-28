package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	findPrime(1000)
	fmt.Println("====")
	filter(1000)
	fmt.Println("+++")
	Find(1000)
}

func filter(n int) {
	total := 0
	filters := make([]int, 0)
	for i := 2; i <= n; i++ {
		filters = append(filters, i)
	}
	primes := make([]int, 0)
	for {
		if len(filters) == 0 {
			break
		}
		prime := filters[0]
		primes = append(primes, prime)
		if len(filters) == 0 {
			break
		}
		filters = filters[1:]
		i := 0
		for _, data := range filters {
			total++
			if data%prime != 0 {
				filters[i] = data
				i++
			}
		}
		filters = filters[:i]
	}
	fmt.Println(primes)
	fmt.Println(total)
}

var totalA uint64

func Find(n int) {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < n; num++ {
		origin <- num
	}
	close(origin)
	<-wait
	fmt.Println(totalA)
}

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			atomic.AddUint64(&totalA, 1)
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}

func findPrime(n int) {
	total := 0
	lists := []int{2}
	for i := 3; i <= n; i++ {
		isPrime := true
		for _, p := range lists {
			total++
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			lists = append(lists, i)
		}
	}
	fmt.Println(lists)
	fmt.Println(total)
}
