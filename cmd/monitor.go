/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:     "monitor",
	Aliases: []string{"monit"},
	Short:   "Start monitoring configured urls",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("config")
		fmt.Printf("%#v", filename)

		var configs = RemoteConfigs{}
		configs.load(filename)

		if len(configs) == 0 {
			log.Fatal("Configuration is empty!")
		}

		fmt.Printf("%-40s%v\n", "URL", "INTERVAL")
		for _, config := range configs {
			fmt.Printf("%-40s%v\n", config.Url, config.Interval)
		}
		fmt.Println()

		fmt.Printf("%-40s%v\n", "URL", "STATUS")
		client := http.Client{}
		for _, config := range configs {
			go monitor(&client, config)
		}

		for {
			time.Sleep(time.Minute)
		}
	},
}

func monitor(client *http.Client, config RemoteConfig) {
	duration, err := time.ParseDuration(config.Interval)
	if err != nil {
		log.Fatalf("Unable to parse interval duration %#v of #%v!\n", config.Interval, config.Url)
	}

	req, err := http.NewRequest("GET", config.Url, nil)
	if err != nil {
		log.Fatalf("Failed to create request for #%v!\n", config.Url)
	}

	for {
		res, err := client.Do(req)
		if err != nil {
			log.Fatalf("Received error: %v\n", err.Error())
		}

		// if res.StatusCode >= 300 {
		fmt.Printf("%-40s%v\n", config.Url, res.Status)
		// }
		time.Sleep(duration)
	}
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
