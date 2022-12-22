package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool ...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests ...")
	result := m.Run()

	fmt.Println("Cleaning up ...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task1 := "enable adding task from STDIN"
	task2 := "enable configuration with env vars"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		// first, test ./todo -task <this is a todo item>
		cmd1 := exec.Command(cmdPath, "-task", task1)

		if err := cmd1.Run(); err != nil {
			t.Fatal(err)
		}

		// next, test. ./todo -add <this is another todo item>
		cmd2 := exec.Command(cmdPath, "-add", task2)
		if err := cmd2.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("  1: %s\n  2: %s\n", task1, task2)

		if string(out) != expected {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		_, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// after completion, list should be empty
		cmd2 := exec.Command(cmdPath, "-list")
		out, err := cmd2.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("x 1: %s\n  2: %s\n", task1, task2)

		if string(out) != expected {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})
}
