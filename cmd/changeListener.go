package cmd

import (
	"fmt"
	"os/exec"
)

type Observer interface {
	Notify(name string, isRunning bool)
}

type MyObserver struct{}

func (o MyObserver) Notify(name string, isRunning bool) {
	var state string
	if isRunning {
		state = "is running"
	} else {
		state = "has stopped"
	}
	displayText := fmt.Sprintf("Observer notified: Service %s %s", name, state)
	execNotify := exec.Command("osascript", "-e", `display notification "`+displayText+`" with title "`+name+`"`)
	_, err := execNotify.Output()
	if err != nil {
		panic("Notify not possible")
	}
}
