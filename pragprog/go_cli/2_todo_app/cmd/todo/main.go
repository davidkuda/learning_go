package main

import (
	"bufio" // read data from the STDIN input stream
	"flag"
	"fmt"
	// io.Reader interface to define param types: decouple your implementations
	// from specific types, allowing your code to work with any types that
	// implement the io.Reader interface
	"io"
	"os"
	"strings"

	"pragprog.com/rggo/interacting/todo"
)

// Default file name, can be overwritten with env var TODO_FILENAME
var todoFileName = ".todo.json"

func main() {
	// Invoke help with custom message with "./todo -h / -help / --help"
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for The Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2020")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information:")
		flag.PrintDefaults()
	}
	// Parse command line flags:
	// we actually no longer need -task, but I keep it as a reference to flag.String()
	task := flag.String("task", "", "Task to be included in the Todo list")
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	// Define an items list
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	l := &todo.List{}

	// Read existing data from .todo.json with Get()
	if err := l.Get(todoFileName); err != nil {
		// exit 1 will facilitate the use of this program in other unix scripts or progs
		// os.Stderr is the place to output err, instead of stdout
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		// List current to do items
		fmt.Print(l)
	case *add:
		// when any arguments (excluding flags) are provided,
		// they will be used as the new task
		// note how os.Stdin matches the required io.Reader interface
		// the ... operator expands the slice into a list of values
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		//Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// Add the task
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Please provide an option with the todo cmd, see todo -help")
		os.Exit(1)
	}
}

// getTask decides where to get the description for a new
// task from: arguments or STDIN
// "variadic function" with three dots as in "parameter ...type"
// -> fn accepts zero or more args
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
		return "", fmt.Errorf("task cannot be blank")
	}
	return s.Text(), nil
}
