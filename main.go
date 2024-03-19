package main

import (
	"aquaminder/notification"
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func welcomeMessage() {
	fmt.Println(Banner)
	fmt.Println()
	fmt.Println("Welcome to AquaMinder and congratulations for beginning your quest of staying hydrated!")
	fmt.Println()
	fmt.Println("You can stop receiving notifications at any time by pressing 'q' to exit the program.")
}

func handleInput(quit chan bool) {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if char == 'q' {
			quit <- true
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	quit := make(chan bool)
	go handleInput(quit)

	welcomeMessage()

	for {
		notification.Notify()

		select {
		case <-time.After(15 * time.Minute):
			// do nothing
		case <-quit:
			// exit in case of quit signal
			return
		}
	}
}
