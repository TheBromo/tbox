package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func processInput(arg string) error {

	subcommand := os.Args[1]

	time, err := strconv.Atoi(subcommand)

	if err != nil {
		return err
	}

	fmt.Printf("⏰ timebox created for %d minutes.\n", time)
	startTimer(time)

	return nil
}

func startTimer(timeMinutes int) {

	time.Sleep(time.Duration(timeMinutes) * time.Minute)

	err := beeep.Notify("Ring ring!!!!", "Your timer of "+strconv.Itoa(timeMinutes)+" min is over", "assets/icon.png")
	if err != nil {
		panic(err)
	}
}

func main() {
	if os.Args[1] == "" {
		processInput("15")
	}
	if err := processInput(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
