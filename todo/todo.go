package todo

import (
	"encoding/json"
	"fmt"
	"os"
)
type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)

	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	f, err := os.ReadFile(filename)

	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(f, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}