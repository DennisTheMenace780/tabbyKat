package main

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list     list.Model
	choice   string
	err      string
	quitting bool
}

type checkoutMsg struct{ err error }

func checkout(branch string) tea.Cmd {
	c := exec.Command("git", "checkout", branch)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return checkoutMsg{err}
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				branch := strings.TrimLeft(string(i), "*")
				branch = strings.TrimSpace(branch)
				m.choice = branch
			}
			return m, checkout(m.choice)
		}

	case checkoutMsg:
		if msg.err != nil {
			return m, tea.Quit
		}
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	// if m.err != "" {
	// 	return QuitCheckoutStyle.Render(m.err)
	// }
	// if m.choice != "" {
	// 	return QuitCheckoutStyle.Render(fmt.Sprintf("Switching to branch ... '%s'", m.choice))
	// }
	// if m.quitting {
	// 	return QuitTextStyle.Render("Not hungry? That’s cool.")
	// }
	return "\n" + m.list.View()
}
