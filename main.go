package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func processInput(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a time")
	}

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
	if err := processInput(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
