package main

import (
	"fmt"
	"net/http"
	"os"
)

var diningHallHost = "http://localhost"
var receivedOrder Order
var orderList OrderList

var cookArr = cookStaff

const stoveNum = 2
const ovenNum = 2

func main() {

	args := os.Args

	if len(args) > 1 {
		//Set the docker internal host
		diningHallHost = args[1]
	}

	fmt.Println("Kitchen is up and running")
	http.HandleFunc("/", responseHandler)
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/send", sendHandler)

	for _, cook := range cookArr {
		go cook.work()
	}
	http.ListenAndServe(":8000", nil)

}
