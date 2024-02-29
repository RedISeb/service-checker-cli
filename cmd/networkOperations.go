package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func ReadDefaultConfig() Services {
	data, err := os.Open("./config/default.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer data.Close()

	byteValue, _ := io.ReadAll(io.Reader(data))
	var services Services
	json.Unmarshal(byteValue, &services)

	return services
}

func PingHost(hostname string) bool {
	app := "ping"
	arg1 := "-c"
	arg2 := "1"
	arg3 := hostname

	execPing := exec.Command(app, arg1, arg2, arg3)
	_, err := execPing.Output()
	return err == nil
}

func (s *Service) CheckTcpPortUnix() {
	app := "nc"
	arg1 := "-vz"
	arg2 := s.Host
	arg3 := s.Port

	execTest := exec.Command(app, arg1, arg2, string(arg3))
	_, err := execTest.Output()
	if err == nil {
		s.SetIsRunning(true)
	}
}

func (s *Service) OutputDialogTCP() {
	switch s.IsRunning {
	case true:
		fmt.Printf("Service on port %s is running.\n", s.Port)
	case false:
		fmt.Printf("Service on port %s is down.\n", s.Port)
	}
}

func GetSystemInfo() string {
	os := runtime.GOOS
	return os
}
