package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var diningHallHost = "http://localhost"



var receivedOrder Order

//type Order struct {
//	OrderId  int `json:"order_id"`
//	TableId  int `json:"table_id"`
//	WaiterId int `json:"waiter_id"`
//	Items    []int `json:"items"`
//	Priority   int   `json:"priority"`
//	MaxWait    int   `json:"max_wait"`
//	PickUpTime int64 `json:"pick_up_time"`
//}

func responseHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "Last received order is:", receivedOrder)
}

var s = "asdasd"
func sendHandler(w http.ResponseWriter, r *http.Request) {
	request,_ := http.NewRequest(http.MethodPost, diningHallHost+":8001", bytes.NewBuffer([]byte(s)))
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Fprintln(w,"ERROR:",err)
	} else {

		fmt.Fprintln(w, "Sent:"+s)

		var responseBuffer = make([]byte,response.ContentLength)

		response.Body.Read(responseBuffer)

		fmt.Fprintln(w, "Received:"+string(responseBuffer))
	}
}

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


	http.ListenAndServe(":8000",nil)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {

	var responseBuffer = make([]byte, r.ContentLength)
	r.Body.Read(responseBuffer)
	var order Order
	json.Unmarshal(responseBuffer, &order)
	fmt.Fprintln(w,order)
	orderList.orderArr = append(orderList.orderArr, &order)


}

var orderList OrderList