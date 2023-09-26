package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type RemoteConfig struct {
	Url      string `json:"url"`
	Interval string `json:"interval"`
}
type RemoteConfigs []RemoteConfig

func (configs *RemoteConfigs) add(config RemoteConfig) {
	*configs = append(*configs, config)
}

func (configs *RemoteConfigs) remove(url string) {
	found := -1
	for index, config := range *configs {
		if config.Url == url {
			found = index
			break
		}
	}

	if found >= 0 {
		*configs = append((*configs)[:found], (*configs)[found+1:]...)
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
