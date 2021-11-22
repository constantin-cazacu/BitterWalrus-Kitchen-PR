package main

import "sync"

type Order struct {
	OrderId        int   `json:"order_id"`
	TableId        int   `json:"table_id"`
	WaiterId       int   `json:"waiter_id"`
	Items          []int `json:"items"`
	Priority       int   `json:"priority"`
	MaxWait        int   `json:"max_wait"`
	PickUpTime     int64 `json:"pick_up_time"`
	CookingDetails []FoodDelivery
}

type OrderList struct {
	orderArr []*Order
	mx       sync.Mutex
}
