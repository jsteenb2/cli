package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jsteenb2/cli/pkg/todo"
)

func main() {
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	var todoFileName = ".todo.json"
	if filename := os.Getenv("TODO_FILENAME"); filename != "" {
		todoFileName = filename
	}

	l := new(todo.List)
	errOut(l.Get(todoFileName))

	switch {
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		errOut(err)
		l.Add(t)
		errOut(l.Save(todoFileName))
	case *list:
		fmt.Print(l)
	case *complete > 0:
		errOut(l.Complete(*complete))
		errOut(l.Save(todoFileName))
	default:
		errOut(errors.New("invalid option"))
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}
	return s.Text(), nil
}

func errOut(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
