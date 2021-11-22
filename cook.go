package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type Cook struct {
	id          int
	rank        int
	proficiency int
	name        string
	catchPhrase string
}

func (c Cook) work() {
	for {
		var foodIndex = -1
		var orderId = -1
		var minWait = math.MaxInt32
		orderList.mx.Lock()
		var now = time.Now().Unix()
		for i, order := range orderList.orderArr {
			for j, item := range order.Items {
				if menu[item].complexity > c.rank {
					continue
				}
				var estimatedTime = order.MaxWait - int(now - order.PickUpTime) - menu[item].prepTime
				if estimatedTime < minWait {
					minWait = estimatedTime
					orderId = i
					foodIndex = j
				}
			}
		}
		if foodIndex != -1 && orderId != -1 {
			var order = orderList.orderArr[orderId]
			order.CookingDetails = append(order.CookingDetails, FoodDelivery{order.Items[foodIndex], c.id})
			var foodId = order.Items[foodIndex]
			fmt.Println("ready item: ", foodId," complexity: ", menu[foodId].complexity , " by ", c.id, " from order: ", order.OrderId)
			order.Items = append(order.Items[:foodIndex], order.Items[foodIndex+1:]...)
			var reqApparatus = menu[foodId].cookingApparatus
			var apparatusArr *[]cookingApparatus = nil
			switch reqApparatus {
			case "oven":
				apparatusArr = &ovenList
			case "stove":
				apparatusArr = &stoveList
			}
			var foundApparatus = false
			var apparatus *cookingApparatus = nil

			if apparatusArr != nil{
				for i, _ := range *apparatusArr {
					apparatus = &(*apparatusArr)[i]
					if apparatus.busyFlag != true {
						apparatus.busyFlag = true
						foundApparatus = true
						break
					}
				}
				if !foundApparatus {
					apparatus = &(*apparatusArr)[rand.Intn(len(*apparatusArr))]
					if apparatus.busyFlag != true {
						apparatus.busyFlag = true
						foundApparatus = true
					}
				}
			}
			orderList.mx.Unlock()
			if foundApparatus == true {
				apparatus.prepareMutex.Lock()
			}


			time.Sleep(time.Second)
			//time.Sleep(time.Second * time.Duration(order.Items[foodIndex]))

			orderList.mx.Lock()
			if foundApparatus == true {
				apparatus.busyFlag = false
				apparatus.prepareMutex.Unlock()
			}
			orderList.mx.Unlock()
		} else {
			orderList.mx.Unlock()
		}
		orderList.mx.Lock()
		for i, order := range orderList.orderArr {
			if len(order.Items) == 0 {
				completedDelivery := newDelivery(order)
				orderStatus := deliver(completedDelivery)
				if orderStatus != false {
					orderList.orderArr = append(orderList.orderArr[:i], orderList.orderArr[i+1:]...)
					fmt.Println("Delivery status: complete --order ", completedDelivery.OrderId)
					break
				}
			}
		}
		orderList.mx.Unlock()
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

var cookStaff = []Cook{{
	id:          0,
	rank:        3,
	proficiency: 3,
	name:        "Master",
	catchPhrase: "yes",
}, {
	id:          1,
	rank:        1,
	proficiency: 1,
	name:        "Jeremy Clarkson",
	catchPhrase: "Look at that! It's Brrrrilliant",
}, {
	id:          2,
	rank:        1,
	proficiency: 3,
	name:        "Lewis Hamilton",
	catchPhrase: "My tires are dead man",
}, {
	id:          3,
	rank:        2,
	proficiency: 2,
	name:        "Cillian Murphy",
	catchPhrase: "I think, that's what I do, so that you don't have to",
}, {
	id:          4,
	rank:        3,
	proficiency: 2,
	name:        "Rowan Atkinson",
	catchPhrase: "*grunts*",
},
}
