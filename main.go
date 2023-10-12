package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"github.com/gen2brain/beeep"
)

func processInput(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a time")
	}

	subcommand := os.Args[1]

	time, err := strconv.Atoi(subcommand)

	if err != nil {
		return err
	}

	fmt.Printf("⏰ timebox created for %d minutes.\n", time)
	go startTimer(time)

	return nil
}

func startTimer(timeMinutes int) {

	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		panic(err)
	}
}

func main() {
	if err := processInput(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
