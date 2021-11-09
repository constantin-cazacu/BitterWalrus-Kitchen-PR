package main

import (
	"fmt"
	"net/http"
	"os"
)

var diningHallHost = "http://localhost"
var receivedOrder Order
var orderList OrderList

var cookTest Cook

func main() {

	args := os.Args

	if len(args) > 1{
		//Set the docker internal host
		diningHallHost = args[1]
	}

	fmt.Println("hello")
	http.HandleFunc("/",responseHandler)
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/send",sendHandler)

	go cookTest.work()

	http.ListenAndServe(":8000",nil)


}



