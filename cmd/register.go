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
	Short: "Register lets you add a new service to be monitored.",
	Long: `Register adds new services to the default.json config file.
Everything that is registered in default.json can be monitored using the monitor command.
The host flag is mandatory, whereas name defaults to the host and the port to 80 if left blank.
Usage:
service-checker-cli register --host localhost --name localhost --port 443
or
service-checker-cli register -H localhost -n localhost -p 443
`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = host
		}
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
	registerCmd.PersistentFlags().StringP("host", "H", "localhost", "Hostname or ip address of the service.")
	registerCmd.MarkPersistentFlagRequired("host")
	registerCmd.PersistentFlags().StringP("name", "n", "", "Name for the registered service.")
	registerCmd.PersistentFlags().StringP("port", "p", "80", "Specify the TCP port of the registered service.")
}
