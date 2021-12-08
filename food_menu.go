package main

type Menu struct {
	id	 			 int
	name	 		 string
	prepTime	 	 int
	complexity	 	 int
	cookingApparatus string
}

var pizza = Menu {1, "Pizza", 20, 2, "oven"}
var salad = Menu {2, "Salad", 10, 1, ""}
var zeama = Menu {3, "zeama", 7, 1, "stove"}
var sashimi = Menu {4, "Scallop Sashimi with Meyer Lemon Confit", 32, 3, ""}
var duck = Menu {5, "Island Duck with Mulberry Mustard", 35, 3, "oven"}
var waffles = Menu {6, "Waffles", 10, 1, "stove"}
var aubergine = Menu {7, "Aubergine", 20, 2, ""}
var lasagna = Menu {8, "Lasagna", 30, 2, "oven"}
var burger = Menu {9, "Burger", 15, 1, "oven"}
var gyros = Menu {10, "Gyros", 15, 1, ""}

var menu = []Menu{pizza, pizza, salad, zeama, sashimi, duck, waffles, aubergine, lasagna, burger, gyros}

