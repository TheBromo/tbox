package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func main(){	
	time , err:= strconv.Atoi(os.Args[1])

	if err != nil{
		fmt.Println("invalid argument")
		return
	}
	startTimer(time)
}



func startTimer(timeMinutes int) {

	time.Sleep(time.Duration(timeMinutes) * time.Minute)

	err := beeep.Notify("Ring ring!!!!", "Your timer of "+strconv.Itoa(timeMinutes)+" min is over", "assets/icon.png")
	if err != nil {
		panic(err)
	}
}
