package tasks

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gandarfh/calendar/calendar/pkg/constants"
	"github.com/gandarfh/calendar/calendar/pkg/styles"
)

func Width() int {
	return ((constants.WindowSize.Width - (constants.WindowSize.Width / 4)) / 7) * 7
}

type active bool

type Model struct {
	active active
}

func (m Model) SetActive(value bool) tea.Cmd {
	return func() tea.Msg {
		m.active = active(value)
		return active(value)
	}
}

func New() tea.Model {
	return Model{active: false}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		// cmd  tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg

	case active:
		m.active = msg
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {

	tasks := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Foreground(styles.Theme.Yellow).Bold(true).Render("Tasks"),
		"",
		"- Task 1",
		"- Task 2",
		"- Task 3",
		"- Task 4",
		"- Task 5",
		"- Task 6",
	)

	if m.active {
		return styles.Container.Tasks.
			Width(constants.WindowSize.Width - Width() - 9).
			Height(((constants.WindowSize.Height-constants.WindowSize.Height/6)/6)*6 + 3).
			BorderForeground(styles.Theme.Yellow).
			Render(tasks)
	} else {
		return styles.Container.Tasks.
			Width(constants.WindowSize.Width - Width() - 9).
			Height(((constants.WindowSize.Height-constants.WindowSize.Height/6)/6)*6 + 3).
			BorderForeground(styles.Theme.Foreground).
			Render(tasks)
	}
}
