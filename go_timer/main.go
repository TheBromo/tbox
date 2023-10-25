package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	time, err := strconv.Atoi(os.Args[1])

	note := os.Args[2]

	if err != nil {
		fmt.Println("invalid argument")
		return
	}
	startTimer(time, note)
}

func startTimer(timeMinutes int, note string) {
	time.Sleep(time.Duration(timeMinutes) * time.Second)
	e, err := os.Executable()
	if err != nil {
		fmt.Println("couldn't get current line e")
		return
	}

	path := path.Dir(e)

	if note == "" {
		err := beeep.Notify("Ring ring!!!!", "Your timer of "+strconv.Itoa(timeMinutes)+" minutes is over ", path+"/assets/icon.png")
		if err != nil {
			panic(err)
		}
	} else {
		err := beeep.Notify("Ring ring!!!", note+"has passed after "+strconv.Itoa(timeMinutes)+" minutes", path+"/assets/icon.png")
		if err != nil {
			panic(err)
		}
	}
}
