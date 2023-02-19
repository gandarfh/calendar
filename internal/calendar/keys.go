package calendar

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit      key.Binding
	NextMonth key.Binding
	PrevMonth key.Binding
	NextYear  key.Binding
	PrevYear  key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return nil
}

var keys = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	NextMonth: key.NewBinding(
		key.WithKeys("j"),
		key.WithHelp("j", "Next Month"),
	),
	PrevMonth: key.NewBinding(
		key.WithKeys("k"),
		key.WithHelp("k", "Prev Month"),
	),

	NextYear: key.NewBinding(
		key.WithKeys("J"),
		key.WithHelp("J", "Next Year"),
	),
	PrevYear: key.NewBinding(
		key.WithKeys("K"),
		key.WithHelp("K", "Prev Year"),
	),
}
