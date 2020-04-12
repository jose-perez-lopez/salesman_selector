package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	salesMen := os.Args[1:]
	salesmanSelected, err := getRandomSalesMan(salesMen)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\a\v\tAn the winer is: %s\n\n", salesmanSelected)
}

func getRandomSalesMan(salesMen []string) (string, error) {
	if len(salesMen) == 0 {
		return "", errors.New("No salesmen provides, someone has to play!")
	}
	numberOfSales := len(salesMen)
	return salesMen[getRandomNumber(numberOfSales)], nil
}

func getRandomNumber(top int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(top))

}
