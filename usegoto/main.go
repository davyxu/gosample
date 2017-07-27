package main

import (
	"errors"
	"fmt"
)

func recvData() error {

	return errors.New("recv error")
}

func closeData() {

	fmt.Println("closeData")
}

func nousegoto() {

	err := recvData()

	if err != nil {
		fmt.Println(err)
		closeData()
		return
	}

	err = recvData()

	if err != nil {
		fmt.Println(err)
		closeData()
		return
	}
}

func usegoto() {

	err := recvData()

	if err != nil {
		goto onErr
	}

	err = recvData()

	if err != nil {
		goto onErr
	}

	return
onErr:

	fmt.Println(err)
	closeData()
}

func main() {

	nousegoto()

	usegoto()

}
