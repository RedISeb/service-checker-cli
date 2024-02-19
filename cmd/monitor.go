/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

type Services struct {
	Services []Service `json:"services"`
}

type Service struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
}

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "The monitor command allows you to automatically check you registered services.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		services := ReadDefaultConfig()
		for _, service := range services.Services {
			isAlive := pingHost(service.Host)
			if isAlive {
				fmt.Printf("%s is up and running \n", service.Host)
				var osSystem string = getSystemInfo()
				switch osSystem {
				case "darwin":
					outputDialogTCP(service.Port, checkTcpPortUnix(service.Host, service.Port))

				case "linux":
					log.Println("Linux")
				default:
					fmt.Printf("%s.\n", osSystem)
				}
			} else {
				panic("Host is not alive")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

func getSystemInfo() string {
	os := runtime.GOOS
	return os
}
