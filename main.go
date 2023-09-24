package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var configFilename = "settings.json"

func main() {
	configs := RemoteConfigs{}
	configs.load(configFilename)
	// load(&configs)

	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)

	addUrl := addCmd.String("url", "", "Url to monitor")
	addInterval := addCmd.String("interval", "", "Interval at which to monitor")

	rmCmd := flag.NewFlagSet("rm", flag.ContinueOnError)
	rmUrl := rmCmd.String("url", "", "Url to monitor")

	if len(os.Args) > 2 {
		addCmd.Parse(os.Args[2:])
		rmCmd.Parse(os.Args[2:])
	}

	client := http.Client{}
	if *addUrl != "" && *addInterval != "" {
		fmt.Printf("%#v\n", []string{*addUrl, *addInterval})
		configs.add(RemoteConfig{Url: *addUrl, Interval: *addInterval})
		configs.save(configFilename)
	}

	if *rmUrl != "" {
		configs.remove(*rmUrl)
		configs.save(configFilename)
	}

	if len(configs) == 0 {
		log.Fatal("Configuration is empty!")
	}

	fmt.Printf("%-40s%v\n", "URL", "INTERVAL")
	for _, config := range configs {
		fmt.Printf("%-40s%v\n", config.Url, config.Interval)
	}
	fmt.Println()

	fmt.Printf("%-40s%v\n", "URL", "STATUS")
	for _, config := range configs {
		go monitor(&client, config)
	}

	for {
		time.Sleep(time.Minute)
	}
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
