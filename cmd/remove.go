/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a configured url",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		fileName, _ := cmd.Flags().GetString("config")
		configName, _ := cmd.Flags().GetString("name")
		removeIndex := -1
		// fmt.Printf("%#v %#v", url, filename)

		var configs = RemoteConfigs{}
		configs.load(fileName)
		if len(url) > 0 {
			removeIndex = configs.indexOfUrl(url)
		}

		if len(configName) > 0 {
			removeIndex = configs.indexOfName(configName)
		}

		configs.removeAt(removeIndex)
		configs.save(fileName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
