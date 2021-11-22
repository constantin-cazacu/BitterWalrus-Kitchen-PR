package main

import "sync"

type cookingApparatus struct {
	busyFlag     bool
	prepareMutex sync.Mutex
}

var ovenList  = make([]cookingApparatus,ovenNum)
var stoveList = make([]cookingApparatus,stoveNum)
