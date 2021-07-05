package todo

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// Item represents a single todo.
type Item struct {
	Priority int
	Text     string
	position int
	Done     bool
}

// ByPriority implements sort.interface for []Item based on
// the Priority & position fields.
type ByPriority []Item

// Label returns the position of this todo item in the list,
// with a period attached.
func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

// SetPriority changes the priority of this todo item.
// [high=1, normal=2].
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

// PrintPriority returns an icon representing the priority
// of this todo item.
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

// PrintDone returns an icon representing the completion state
// of this todo item.
func (i *Item) PrintDone() string {
	if i.Done {
		// return "☑️"
		return "✔"
	}
	return ""
}

func (items ByPriority) Len() int {
	return len(items)
}

func (items ByPriority) Less(i, j int) bool {
	// make sure completed items are under in the list.
	if items[i].Done != items[j].Done {
		return items[j].Done
	}

	if items[i].Priority != items[j].Priority {
		return items[i].Priority < items[j].Priority
	}

	return items[i].position < items[j].position
}

func (items ByPriority) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// SaveItems stores todo items to a given file.
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

// ReadItems retrieves todo items from a given file and returns it.
func ReadItems(filename string, showFileErr bool) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		// when no need to show error if file does not exist.
		if !showFileErr && strings.HasSuffix(err.Error(), "no such file or directory") {
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
