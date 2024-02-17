package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

var Port string = "80"

func main() {
	hostPointer := flag.String("host", "localhost", "Hostname or IP address where the service runs")
	flag.Parse()
	host := *hostPointer
	fmt.Printf("Check the host: %s\n", host)
	isAlive := pingHost(host)
	fmt.Printf("%s is up and running \n", host)
	if isAlive {
		var osSystem string = getSystemInfo()
		switch osSystem {
		case "darwin":
			log.Println("Mac OS")
			outputDialogTCP(Port, checkTcpPortUnix(host, Port))

		case "linux":
			log.Println("Linux")
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

func checkTcpPortUnix(hostname string, port string) bool {
	app := "nc"
	arg1 := "-vz"
	arg2 := hostname
	arg3 := port

	execTest := exec.Command(app, arg1, arg2, string(arg3))
	_, err := execTest.Output()
	return err == nil
}

func outputDialogTCP(port string, isOpenTCP bool) {
	switch isOpenTCP {
	case true:
		fmt.Printf("Service on port %s is running.\n", port)
	case false:
		fmt.Printf("Service on port %s is down.\n", port)
	}
}
