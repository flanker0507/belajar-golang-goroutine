package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")

	time.Sleep(1 * time.Second)

}

func TestRunAsynchronous(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("D O N E")
}
