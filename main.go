package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	hostPointer := flag.String("host", "localhost", "Hostname or IP address where the service runs")
	flag.Parse()
	host := *hostPointer
	fmt.Printf("Check the host: %s\n", host)
	isAlive := pingHost(host)
	if isAlive {
		var osSystem string = getSystemInfo()
		switch osSystem {
		case "darwin":
			fmt.Println("Mac OS")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s.\n", osSystem)
		}
	} else {
		panic("Host is not alive")
	}
}

func getSystemInfo() string {
	os := runtime.GOOS
	return os
}

func pingHost(hostname string) bool {
	app := "ping"
	arg1 := "-c"
	arg2 := "1"
	arg3 := hostname

	execPing := exec.Command(app, arg1, arg2, arg3)
	_, err := execPing.Output()
	return err == nil
}
