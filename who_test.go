package main

import (
	"testing"
)

func TestGetRandomNumber(t *testing.T) {
	top := 10
	repetitions := top * 10
	var frecuencies = make([]int, top)

	for i := 0; i < repetitions; i++ {
		r := getRandomNumber(top)
		frecuencies[r]++
	}
	sumFrecuency := 0
	for i := 0; i < top; i++ {
		sumFrecuency += frecuencies[i]
	}
	/**
	for i := 0; i < top; i++ {
		log.Printf("Frecuency of [%d] %d%% time", i, frecuencies[i]*100/sumFrecuency)
	}**/
	for i := 0; i < top; i++ {
		if frecuencies[i] == 0 {
			t.Errorf("%d din´t happend anytime, it is posible but smell bad", i)
		}
	}
	if sumFrecuency != repetitions {
		t.Errorf("frecuency sum %d, diferent of repetitions %d", sumFrecuency, repetitions)
	}
}

func TestGetRandomSalesMan(t *testing.T) {
	salesMen := []string{"Chris", "Pepe", "Nick"}
	repetitions := len(salesMen) * 10
	var frecuencies = make(map[string]int)
	for i := 0; i < repetitions; i++ {
		salesMan, _ := getRandomSalesMan(salesMen)
		if frecuencies[salesMan] == 0 {
			frecuencies[salesMan] = 1
		} else {
			frecuencies[salesMan]++
		}
	}
	for _, name := range salesMen {
		//log.Printf("%s -> %d\n", name, frecuencies[name])
		if frecuencies[name] == 0 {
			t.Errorf("%s din´t happend anytime, it is posible but smell bad", name)
		}
	}
}

func TestGetRandomSalesManWithNoNames(t *testing.T) {
	var names = []string{}
	_, err := getRandomSalesMan(names)
	if err == nil {
		t.Errorf("I call getRandomSalesMan with no names and it did report an error")
	}
}
