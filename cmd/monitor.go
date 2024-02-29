/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

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
			isAlive := PingHost(service.Host)
			if isAlive {
				fmt.Printf("%s is up and running \n", service.Host)
				service.CheckTcpPortUnix()
				var osSystem string = GetSystemInfo()
				switch osSystem {
				case "darwin":
					service.OutputDialogTCP()
				case "linux":
					log.Println("Linux")
					service.OutputDialogTCP()
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
	monitorCmd.Flags().BoolP("notify", "n", false, "Toggle system notification if status of service is changing.")
}
