/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		s := &Service{Name: host, Host: host, Port: port, IsRunning: false}
		observer := MyObserver{}
		s.RegisterObserver(observer)
		os := GetSystemInfo()
		isHostAlive := PingHost(host)
		if isHostAlive {
			switch os {
			case "darwin", "linux":
				s.CheckTcpPortUnix()
				s.OutputDialogTCP()
			case "windows":
				fmt.Println("To be implemented")
			}
		} else {
			fmt.Printf("Host %s is down", host)
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.PersistentFlags().StringP("host", "H", "localhost", "Hostname of the system you want to check")
	checkCmd.PersistentFlags().StringP("port", "p", "80", "Specify the TCP port of the service.")
	checkCmd.MarkPersistentFlagRequired("host")
}
