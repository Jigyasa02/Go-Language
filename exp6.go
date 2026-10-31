package main

import (
	"fmt"
	"sync"
)

func calculateSquares(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Square Goroutine Started")

	for i := 1; i <= 5; i++ {
		result := i * i
		ch <- fmt.Sprintf("Square of %d = %d", i, result)
	}

	fmt.Println("Square Goroutine Completed")
}

func calculateCubes(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Cube Goroutine Started")

	for i := 1; i <= 5; i++ {
		result := i * i * i
		ch <- fmt.Sprintf("Cube of %d = %d", i, result)
	}

	fmt.Println("Cube Goroutine Completed")
}

func calculateFibonacci(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Fibonacci Goroutine Started")

	a, b := 0, 1

	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Fibonacci number %d = %d", i, a)
		a, b = b, a+b
	}

	fmt.Println("Fibonacci Goroutine Completed")
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(3)

	go calculateSquares(ch, &wg)
	go calculateCubes(ch, &wg)
	go calculateFibonacci(ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Main function collecting results:")

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("All goroutines completed.")
}
