package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func init() {
    // Will run at startup and log to a file
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
	cmd := exec.Command("git", "branch")
	// Get a pipe to read from standard out
	r, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatalf("Error creating stdout pipe: %s", err)
    }
	cmd.Stderr = cmd.Stdout
	// Make a new channel which will be used to ensure we get all output
	done := make(chan struct{})
	// Create a scanner which scans r in a line-by-line fashion
	scanner := bufio.NewScanner(r)
	// Use the scanner to scan the output line by line and log it
	// It's running in a goroutine so that it doesn't block
	var branches []string
	go func() {
		// Read line by line and process it
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "not a git repository") {
                fmt.Println(line)
                os.Exit(1)
			}
			branches = append(branches, line)
		}
		// We're all done, unblock the channel
		done <- struct{}{}
	}()
	// Start the command and check for errors
	err = cmd.Start()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// Wait for all output to be processed
	<-done

	branchItems := BuildItems(branches)
	if err != nil {
        log.Print("Error: ", err)
	}
	l := ListBuilder(branchItems)

	err = cmd.Wait()
	if err != nil {
		log.Print("Error: ", err)
	}

    p := tea.NewProgram(Model{list: l})
    m, err := p.Run()
    if err != nil {
        log.Fatal(err)
    }
	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(Model); ok && m.choice != "" {
        p.Kill()
		fmt.Printf("\n---\nYou chose %s!\n", m.choice)
	}
}
