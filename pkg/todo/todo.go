package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	*l = append(*l, item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	})
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(l)
}

func (l *List) String() string {
	var sb strings.Builder
	for i, item := range *l {
		prefix := "  "
		if item.Done {
			prefix = "✔ "
		}
		fmt.Fprintf(&sb, "%s%d: %s\n", prefix, i+1, item.Task)
	}
	return sb.String()
}
