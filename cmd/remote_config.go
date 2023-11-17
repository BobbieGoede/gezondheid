package cmd

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Plugin struct {
	Name   string         `json:"name"`
	Config map[string]any `json:"config"`
}

type RemoteConfig struct {
	Name     string   `json:"name"`
	Url      string   `json:"url"`
	Interval string   `json:"interval"`
	Plugins  []Plugin `json:"plugins"`
}
type RemoteConfigs []RemoteConfig

func (configs *RemoteConfigs) add(config RemoteConfig) {
	parsed, err := url.Parse(config.Url)
	if err != nil {
		log.Fatal("Failed to parse url!")
	}

	if len(config.Name) == 0 {
		config.Name = strings.ReplaceAll(parsed.Host, ".", "-")
	}

	*configs = append(*configs, config)
}

func (configs *RemoteConfigs) indexOfUrl(url string) int {
	found := -1
	for index, config := range *configs {
		if config.Url == url {
			found = index
			break
		}
	}

	return found
}

func (configs *RemoteConfigs) indexOfName(name string) int {
	found := -1
	for index, config := range *configs {
		if config.Name == name {
			found = index
			break
		}
	}

	return found
}

func (configs *RemoteConfigs) removeAt(index int) {
	if index >= 0 {
		*configs = append((*configs)[:index], (*configs)[index+1:]...)
	}
}

func (configs *RemoteConfigs) save(filename string) {
	fmt.Printf("%#v\n", configs)
	b, err := yaml.Marshal(configs)
	if err != nil {
		log.Fatalf("Unable to marshal data %v\n", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file %v\n", err)
	}

	n, err := f.Write(b)
	if err != nil {
		log.Fatalf("Failed to write to file %v\n", err)
	}
	defer f.Close()

	fmt.Printf("Wrote %d bytes to file\n", n)
}

func (configs *RemoteConfigs) load(filename string) {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating initial settings file")
		configs.save(filename)
		return
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(b, &configs)
	if err != nil {
		log.Fatal(err)
	}
}
