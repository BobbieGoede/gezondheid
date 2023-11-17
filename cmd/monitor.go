/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/bobbiegoede/gezondheid/internal/handlers"
	"github.com/bobbiegoede/gezondheid/internal/plugins"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
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

type DefaultHandler struct {
	config   RemoteConfig
	client   *http.Client
	req      *http.Request
	duration time.Duration
	next     handlers.Handler
}

func (h *DefaultHandler) HandleRequest(ctx *handlers.Ctx) {
	res, err := h.client.Do(h.req)
	if err != nil {
		log.Fatalf("Received error: %v\n", err.Error())
	}

	ctx.Proto = res.Proto
	ctx.StatusCode = res.StatusCode

	// if res.StatusCode >= 300 {
	fmt.Printf("%-40s%v\n", h.config.Url, res.Status)
}

func (h *DefaultHandler) SetNext(handler handlers.Handler) {
	h.next = handler
}

func monitor(client *http.Client, config RemoteConfig) {
	duration, err := time.ParseDuration(config.Interval)
	if err != nil {
		log.Fatalf("Unable to parse interval duration %#v of %#v!\n", config.Interval, config.Url)
	}

	req, err := http.NewRequest("GET", config.Url, nil)
	if err != nil {
		log.Fatalf("Failed to create request for #%v!\n", config.Url)
	}

	var hs []handlers.Handler

	for _, ps := range config.Plugins {
		p := plugins.LoadPlugin(ps.Name)
		ymlData, err := json.Marshal(ps.Config)
		if err != nil {
			log.Fatalf("Failed to create request for %#v!\n", config.Url)
		}

		p.SetConfig(ymlData)
		h := &plugins.PluginHandler{Plugin: p}
		hs = append(hs, h)
	}

	h := &DefaultHandler{req: req, duration: duration, client: client, config: config}
	hs = append(hs, h)

	var first = handlers.SetNextReferences(hs)

	for {
		first.HandleRequest(&handlers.Ctx{})
		time.Sleep(duration)
	}
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
