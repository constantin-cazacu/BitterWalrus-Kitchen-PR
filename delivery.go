package main

import "time"

type Delivery struct {
	OrderId        int            `json:"order_id"`
	TableId        int            `json:"table_id"`
	Items          []int          `json:"items"`
	Priority       int            `json:"priority"`
	MaxWait        int            `json:"max_wait"`
	PickUpTime     int64          `json:"pick_up_time"`
	CookingTime    int            `json:"cooking_time"`
	CookingDetails []FoodDelivery `json:"cooking_details"`
}

type FoodDelivery struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

func newDelivery(order *Order) *Delivery {
	ret := new(Delivery)
	ret.OrderId = order.OrderId
	ret.TableId = order.TableId
	var itemsArr []int
	for _, detail := range order.CookingDetails {
		itemsArr = append(itemsArr, detail.FoodId)
	}
	ret.Items = itemsArr
	ret.Priority = order.Priority
	ret.MaxWait = order.MaxWait
	ret.PickUpTime = order.PickUpTime
	ret.CookingTime = int(time.Now().Unix() - order.PickUpTime)
	ret.CookingDetails = order.CookingDetails
	return ret
}
