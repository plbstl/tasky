package todo

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Priority int
	Text     string
	position int
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) SetPriority(p int) {
	switch p {
	case 1:
		// high
		i.Priority = 1
	default:
		// normal
		i.Priority = 2
	}
}

func (i *Item) PrintPriority() string {
	switch i.Priority {
	case 1:
		// high
		return "⚠"
	default:
		// normal
		return "\u200a"
	}
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

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		// if file does not exist, it is not an error
		if strings.HasSuffix(err.Error(), "no such file or directory") {
			return []Item{}, nil
		}
		return []Item{}, err
	}

	var items []Item

	if err = json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}
