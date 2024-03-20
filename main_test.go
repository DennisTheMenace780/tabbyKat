package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/muesli/termenv"
)

func init() {
	// This is required for CI to pass. See https://charm.sh/blog/teatest/
	lipgloss.SetColorProfile(termenv.Ascii)
}

func setup() {
	if err := os.Chdir("./TestRepo/"); err != nil {
		log.Fatal("Error changing directory: ", err)
	}
}

func cleanup() {
	resetBranchAfterTest()
}

func TestMain(m *testing.M) {
    setup() 

	exitCode := m.Run()

	cleanup()

	os.Exit(exitCode)
}

func TestOutput(t *testing.T) {
	t.Run("Move down and select a branch", func(t *testing.T) {

		model := initModel([]string{"branch-1", "branch-2", "branch-3"})

		tm := teatest.NewTestModel(t, model, teatest.WithInitialTermSize(300, 100))

		// Assert that the program, at some point, has the following byte string ... make a helper function?
		teatest.WaitFor(t, tm.Output(),
			func(bts []byte) bool {
				return bytes.Contains(
					bts,
					[]byte("1. branch-1"),
				)
			},
		)

		moveDownAndSelectBranch(tm, 1)

		tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second))

		out, err := io.ReadAll(tm.FinalOutput(t))
		if err != nil {
			t.Error("Error reading from FinalOutput", err)
		}
		teatest.RequireEqualOutput(t, out)
	})
}
//
// func tearDownGitTestRepository() {
// 	cmd := exec.Command("rm", "-rf", "TestOutput")
// 	cmd.Dir = "./testdata/"
// 	cmd.Run()
// }

func resetBranchAfterTest() {
	cmd := exec.Command("git", "checkout", "branch-1")
	err := cmd.Run()
	if err != nil {
        log.Fatal("Problem resetting branch: ", err)
	}
}

func initModel(branches []string) Model {
	var items []list.Item
	for _, branch := range branches {
		items = append(items, Item(branch))
	}
	l := list.New(items, ItemDelegate{}, DefaultWidth, ListHeight)
	return Model{list: l}
}

func moveDownAndSelectBranch(tm *teatest.TestModel, down int) {
	for i := 0; i < down; i++ {
		tm.Send(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune("j"),
		})
	}

	tm.Send(tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune("enter"),
	})
}
