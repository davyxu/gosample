package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", ":8801")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer l.Close()

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("accepted!")

		go onRequest(conn)
	}
}

func onRequest(conn net.Conn) {

	for {
		reader := bufio.NewReader(conn)

		data, err := reader.ReadString('\n')

		data = strings.TrimSpace(data)

		if err != nil || data == "quit" {
			fmt.Println("closed!")
			conn.Close()
			return
		}

		fmt.Println(data, []byte(data))

		conn.Write([]byte(data + "\r\n"))
	}

}
