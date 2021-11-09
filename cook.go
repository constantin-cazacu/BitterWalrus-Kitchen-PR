package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Cook struct {
	id 			int
	rank 		int
	proficiency int
	name        string
	catchPhrase string
}

func (c Cook) work() {
	for {
		var foodId = -1
		var deliveryOrderId = -1
		for i, order := range orderList.orderArr {
			if len(order.Items) > 0 {
				foodId = order.Items[0]
				order.Items = append(order.Items[:0], order.Items[1:]...)
			}

			if foodId != -1 {
				//time.Sleep(time.Duration(menu[foodId].prepTime) * time.Second)
				time.Sleep(time.Second)
				order.CookingDetails = append(order.CookingDetails, FoodDelivery{foodId, c.id})
				fmt.Println("ready item: ", foodId, " by ", c.id)
			} else {
				deliveryOrderId = i
			}
			break
		}

		if deliveryOrderId != -1 {
			completedDelivery := newDelivery(orderList.orderArr[deliveryOrderId])
			orderStatus := deliver(completedDelivery)
			if orderStatus != false {
				orderList.orderArr = append(orderList.orderArr[:deliveryOrderId], orderList.orderArr[deliveryOrderId+1:]...)
			}
		}

		time.Sleep(time.Second)
		}
}

func deliver(delivery *Delivery) bool {

	requestBody, marshallErr := json.Marshal(delivery)
	if marshallErr != nil {
		log.Fatal(marshallErr)
	}

	request, newRequestError := http.NewRequest(http.MethodPost, diningHallHost+":8001"+"/delivery", bytes.NewBuffer(requestBody))
	if newRequestError != nil {
		fmt.Println("Could not create new request. Error:", newRequestError)
		log.Fatal(newRequestError)
	} else {
		response, doError := http.DefaultClient.Do(request)
		fmt.Println("Delivery attempt")
		if doError != nil {
			fmt.Println("ERROR Sending request. ERR:", doError)
			log.Fatal(doError)
		}
		var responseBody = make([]byte, response.ContentLength)
		response.Body.Read(responseBody)
		fmt.Println("Response: ", string(responseBody))
		if string(responseBody) != "OK" {
			return false
		}
		return true
	}
	return true
}



var cookStaff = []Cook {{
	rank: 		 1,
	proficiency: 1,
	name: 		 "Jeremy Clarkson",
	catchPhrase: "Look at that! It's Brrrriliant",
}, {
	rank: 		 1,
	proficiency: 3,
	name: 		 "Lewis Hamilton",
	catchPhrase: "My tires are dead man",
}, {
	rank: 		 2,
	proficiency: 2,
	name: 		 "Cillian Murphy",
	catchPhrase: "I think, that's what I do, so that you don't have to",
}, {
	rank: 		 3,
	proficiency: 2,
	name: 		 "Rowan",
	catchPhrase: "grunts",
},
}


