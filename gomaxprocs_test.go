package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalcpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalcpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines:", totalGoroutines)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalcpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalcpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines:", totalGoroutines)

	group.Wait()
}
