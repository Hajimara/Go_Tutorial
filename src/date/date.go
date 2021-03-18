package main

import (
	"fmt"
	"constants"
	"log"
)

func main()  {
	date := constants.Date{}
	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(3)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.GetYear())
	fmt.Println(date.GetMonth())
	fmt.Println(date.GetDay())
	fmt.Println(date.GetDate())
	event := constants.Event{}
	err = event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(3)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.GetYear())
	fmt.Println(event.GetMonth())
	fmt.Println(event.GetDay())
	fmt.Println(event.Date.GetDate())
}