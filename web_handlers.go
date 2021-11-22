package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func responseHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "Last received order is:", receivedOrder)
}


func sendHandler(w http.ResponseWriter, r *http.Request) {
	request,_ := http.NewRequest(http.MethodPost, diningHallHost+":8001", bytes.NewBuffer([]byte("placeholder")))
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Fprintln(w,"ERROR:",err)
	} else {

		fmt.Fprintln(w, "Sent:")

		var responseBuffer = make([]byte,response.ContentLength)

		response.Body.Read(responseBuffer)

		fmt.Fprintln(w, "Received:"+string(responseBuffer))
	}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {

	var responseBuffer = make([]byte, r.ContentLength)
	r.Body.Read(responseBuffer)
	var order Order
	json.Unmarshal(responseBuffer, &order)
	fmt.Fprintln(w,order)
	orderList.orderArr = append(orderList.orderArr, &order)
	fmt.Println("Order received: ", order.OrderId, order.Items, order.MaxWait)

}
