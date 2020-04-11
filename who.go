package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numberOfSales := 10
	salesSelected := rand.Intn(numberOfSales)
	fmt.Println(salesSelected)

}
