package main

import (
	"fmt"
	"math"
)

func main() {
	m := make(map[string]float64)

	//keys and values to the map
	m["Noodles"] = 1.60
	m["Coffee"] = 2.50
	m["Chicken filets"] = 4.55
	m["Wine"] = 4.55
	m["Apples"] = 1.55
	m["Bread"] = 1.95
	m["Pudding"] = 0.95
	m["Avocado"] = 2.05

	//slice of values from the map
	prices := []float64{}
	for _, value := range m {
		prices = append(prices, value)
	}
	//fmt.Printf("%g", prices)

	//slice of keys from the map
	items := []string{}
	for key := range m {
		items = append(items, key)
	}
	fmt.Println("Shopping list", items)

	printMap(m)
	av := sum(prices)

	var tax = 19.00
	tax = tax * av / 100
	tax = math.Round(tax)
	fmt.Println("Tax:", tax, "Euro")
}

//Sum of all prices
func sum(prices []float64) float64 {
	av := 0.00

	for _, value := range prices {
		av += value
	}
	fmt.Println("Total price:", av, "Euro")
	return av
}

//list of all items with prices/map
func printMap(m map[string]float64) {
	for item, price := range m {
		fmt.Println("Price for", item, "is", price)
	}
}
