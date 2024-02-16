package main

import (
	"fmt"
	"runtime"
)

func main() {
	var osSystem string = getSystemInfo()
	fmt.Println(osSystem)
}

func getSystemInfo() string {
	os := runtime.GOOS
	return os
}
