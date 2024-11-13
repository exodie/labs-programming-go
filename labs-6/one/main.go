package one

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Факториал %d = %d\n", n, result)
}

func randomNumbers(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		fmt.Printf("Случайное число: %d\n", num)
		time.Sleep(200 * time.Millisecond)
	}
}

func sumSeries(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Printf("Сумма числового ряда от 1 до %d = %d\n", n, sum)
}

func Run() {
	var wg sync.WaitGroup

	wg.Add(3)

	go factorial(5, &wg)
	go randomNumbers(5, &wg)
	go sumSeries(10, &wg)

	wg.Wait()
}
