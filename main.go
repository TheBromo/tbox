package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
	"os/exec"
)

func processInput(arg string) error {

	subcommand := os.Args[1]

	if subcommand == "" {
		subcommand = "15"
	}

	for _, c := range subcommand {
		if !unicode.IsDigit(c) {
			return errors.New("argument is not a number")	
		}
	}

	fmt.Printf("⏰ timebox created for %s minutes.\n", subcommand)

	startTimerProcess(subcommand)
	return nil
}

func startTimerProcess(time string) error {
	cmd := exec.Command("go", "run", "./go_timer", time)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return errors.New( err.Error())
	}
	err = cmd.Process.Release()
	if err != nil {
		return errors.New("cmd.Process.Release failed: ")
	}
	return nil
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
