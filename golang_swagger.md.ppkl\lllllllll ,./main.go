package main

import (
	"fmt"
	"math/rand"
	"time"
)

var listForks = make([]chan struct{}, 0, 5)
var listPerson = make([]chan struct{}, 0, 5)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		listForks = append(listForks, make(chan struct{}, 1))
		listForks[i] <- struct{}{}
	}

	for i := 0; i < 5; i++ {
		listPerson = append(listPerson, make(chan struct{}, 1))
	}
	go randomGenerateHungry()

	fmt.Println("init randomGenerateHungry")
	go func() {

		for {
			select {

			case <-listPerson[0]:
				reserve(0)

			case <-listPerson[1]:
				reserve(1)

			case <-listPerson[2]:
				reserve(2)

			case <-listPerson[3]:
				reserve(3)

			case <-listPerson[4]:
				reserve(4)
			default:
				fmt.Println("no body si not hungry")
				time.Sleep(time.Second)

			}

		}

	}()

	time.Sleep(time.Second * 100000)

}

func randomGenerateHungry() {
	for {

		randPerson := rand.Intn(5)
		if len(listPerson[randPerson]) == 0 {

			listPerson[randPerson] <- struct{}{}
			fmt.Println("listPerson ", randPerson, "is hungry")

			time.Sleep(3 * time.Second)
		}
	}

}
func reserve(person int) {
	firstFork := person
	secondFork := person + 1
	if secondFork > 4 {

		secondFork = 0
	}

	select {
	case <-listForks[firstFork]:
		<-listForks[secondFork]
		eat(person)
	case <-listForks[secondFork]:
		<-listForks[firstFork]
		eat(person)

	}

}

func eat(person int) {
	firstFork := person
	secondFork := person + 1
	if secondFork > 4 {

		secondFork = 0
	}

	fmt.Println("eating person", person)
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	fmt.Println("finish eat  person", person)
	listForks[firstFork] <- struct{}{}
	listForks[secondFork] <- struct{}{}

	<-listPerson[person]

}
