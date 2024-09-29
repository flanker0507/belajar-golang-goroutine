package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Yuda Wira Pratama"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeRespond(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Yuda Wira Pratama"
}

func TestGiveMeRespond(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespond(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Yuda Wira Pratama"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)

}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)
	go func() {
		channel <- "Yuda"
		channel <- "Wira"
		channel <- "Pratama"

	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("DONE")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke - " + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data", data)
	}

	fmt.Println("DONE")

}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		GiveMeRespond(channel1)
		GiveMeRespond(channel2)
	}()

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1 ", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data Dari Channel 2 ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
	fmt.Println("D O N E")
}

func TestSelectDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		GiveMeRespond(channel1)
		GiveMeRespond(channel2)
	}()

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1 ", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data Dari Channel 2 ", data)
			counter++

		default:
			fmt.Println("Waiting Of Data")
		}

		if counter == 2 {
			break
		}
	}
	fmt.Println("D O N E")
}
