package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(header string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string) {
	for {
		message := <-channel // 此处会阻塞, 如果信道中没有数据的话
		fmt.Println(message)
	}
}

func main() {
	channel := make(chan string)
	go producer("cat", channel)
	go producer("dog", channel)
	customer(channel)
}
