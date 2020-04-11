package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	salesMen := []string{"Chris", "Pepe", "Nick"}
	salesmanSelected := getRandomSalesMan(salesMen)
	fmt.Println(salesmanSelected)
}

func getRandomSalesMan(salesMen []string) string {
	numberOfSales := len(salesMen)
	return salesMen[getRandomNumber(numberOfSales)]
}

func getRandomNumber(top int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(top))

}
