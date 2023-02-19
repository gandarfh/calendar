package internal

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Calendar key.Binding
	Tasks    key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Tasks, k.Calendar}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return nil
}

var keys = KeyMap{
	Calendar: key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "Go to calendar"),
	),
	Tasks: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "Go to tasks"),
	),
}
