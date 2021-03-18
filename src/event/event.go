package main

import (
	"constants"
	"log"
	"fmt"
)
func main()  {
	event := constants.Event{}
	err := event.SetTitle("titlesss3sdfadsfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.GetTitle())
}