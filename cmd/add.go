/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a url to monitor",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		interval, _ := cmd.Flags().GetString("interval")
		filename, _ := cmd.Flags().GetString("config")
		configname, _ := cmd.Flags().GetString("name")

		if len(url) == 0 {
			log.Fatal("Url should not be empty")
		}

		var configs = RemoteConfigs{}
		configs.load(filename)
		configs.add(RemoteConfig{Url: url, Interval: interval, Name: configname})
		configs.save(filename)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("interval", "i", "10m", "Duration of the check interval")
}
