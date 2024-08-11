package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func init() {
    // init functions in Go will run before the main execution thead
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
}

func main() {
	branches, err := getGitBranches()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	branchItems := BuildItems(branches)
	l := BuildListFromItems(branchItems)

	if _, err := tea.NewProgram(Model{list: l}, tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}

func getGitBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	pr, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("error creating stdout pipe: %w", err)
	}
	cmd.Stderr = cmd.Stdout

	if err := cmd.Start(); err != nil { // cmd.Start() spawns a new process
		return nil, fmt.Errorf("error starting command: %w", err)
	}

    // The output of the process from 49 will be streamed to the pipe when it
    // begins to execute, and then we can operate on it concurrently with
    // another go routine.
	branches, err := captureCmdOutput(pr)
	if err != nil {
		return nil, fmt.Errorf("error capturing command output: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("error waiting for command to finish: %w", err)
	}

	return branches, nil
}

func captureCmdOutput(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var branches []string
	done := make(chan error)

    // The goroutine allows the program to create a lightweight thread to parse
    // the output of the command that was written to the io.Reader instance.
    // This approach means that the output of the command will be streamed to
    // the pipe and then the go routine can operate on those lines as they're
    // coming in to ensure we're not blocking the main thread of execution. 
	go func() { 
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "not a git repository") {
				done <- fmt.Errorf(line)
				return
			}
			branches = append(branches, line)
		}
		done <- scanner.Err()
	}()

	if err := <-done; err != nil {
		return nil, fmt.Errorf("error reading command output: %w", err)
	}

	return branches, nil
}
