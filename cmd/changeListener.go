package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

type Observer interface {
	Notify(name string, isRunning bool, isNotify bool)
}

type MyObserver struct{}

func (o MyObserver) Notify(name string, isRunning bool, isNotify bool) {
	var state string
	if isRunning {
		state = "is running"
	} else {
		state = "has stopped"
	}
	if isNotify {
		displayText := fmt.Sprintf("Observer notified: Service %s %s", name, state)
		execNotify := exec.Command("osascript", "-e", `display notification "`+displayText+`" with title "`+name+`"`)
		_, err := execNotify.Output()
		if err != nil {
			panic("Notify not possible")
		}
	} else {
		currentTime := time.Now().Format(time.RFC850)
		switch isRunning {
		case true:
			log.Printf("%s \t %s is running", currentTime, name)
		case false:
			log.Printf("%s \t %s is down", currentTime, name)
		}

	}

}
