// For testing the bubble tea library

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	m := NewModel()
	// New program w/ Initial model and program options
	p := tea.NewProgram(m)

	// Run
	_, err := p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

// Model : App state
type Model struct {
	title string

	textinput textinput.Model
}

// NewModel : Initial Model
func NewModel() Model {

	ti := textinput.New()
	ti.Placeholder = "Enter search term"
	ti.Focus()
	return Model{
		title:     "hello world",
		textinput: ti,
	}
}

// Init : kick off the event loop
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update: handle Msgs
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			v := m.textinput.Value()
			return m, handleQuerySearch(v)
		}
	}

	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

// View: return a string based on the state of our model
func (m Model) View() string {
	s := m.textinput.View()
	return s
}

func handleQuerySearch(query string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("https://api.urbandictionary.com/v0/define?term=%s", url2.QuerEscape(query))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	}
}
