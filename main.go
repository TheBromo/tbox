package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"unicode"
)

func main() {
	if err := processInput(os.Args[1]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func processInput(arg string) error {
	time := "15"
	comment := ""

	if len(os.Args) > 1 {
		time = os.Args[1]
		if len(os.Args) > 2 {
			comment = os.Args[2]
		}
	}

	for _, c := range time {
		if !unicode.IsDigit(c) {
			return errors.New("argument is not a number")
		}
	}

	fmt.Printf("[%s] ⏰ timebox created for %s minutes.\n", comment, time)

	err := startTimerProcess(time, comment)
	return err
}

func startTimerProcess(time string, comment string) error {
	e, err := os.Executable()

	cmd := exec.Command(path.Dir(e)+"/go_timer/go_timer", time, comment)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return errors.New(err.Error())
	}
	err = cmd.Process.Release()
	if err != nil {
		return errors.New("cmd.Process.Release failed: ")
	}
	return nil
}
