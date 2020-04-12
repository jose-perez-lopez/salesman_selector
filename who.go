package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	salesMen := os.Args[1:]
	salesmanSelected := getRandomSalesMan(salesMen)
	fmt.Printf("\a\v\tAn the winer is: %s\n\n", salesmanSelected)
}

func getRandomSalesMan(salesMen []string) string {
	if len(salesMen) == 0 {
		log.Fatal("No salesmen provides, someone has to play!")
	}
	numberOfSales := len(salesMen)
	return salesMen[getRandomNumber(numberOfSales)]
}

func getRandomNumber(top int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(top))

}
