/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		newService := Service{}
		newService.Name = name
		newService.Host = host
		newService.Port = port

		services := ReadDefaultConfig()
		services.Services = append(services.Services, newService)
		updatedData, err := json.MarshalIndent(services, "", "    ")
		if err != nil {
			log.Fatalln(err.Error())
		}

		err = os.WriteFile("./config/default.json", updatedData, 0644)
		if err != nil {
			log.Fatalln(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.PersistentFlags().StringP("host", "H", "localhost", "Hostname or ip address of the service.")
	registerCmd.MarkPersistentFlagRequired("host")
	host, _ := registerCmd.Flags().GetString("host")
	registerCmd.PersistentFlags().StringP("name", "n", host, "Name for the registered service.")
	registerCmd.PersistentFlags().StringP("port", "p", "80", "Specify the TCP port of the registered service.")
}
