package main

import (
	"fmt"
	"os"
	"strings"

	"pragprog.com/rggo/interacting/todo"
)

// Hardcoding the file name
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		// exit 1 will facilitate the use of this program in other unix scripts or progs
		// os.Stderr is the place to output err, instead of stdout
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
		// for no extra args, print the list
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
