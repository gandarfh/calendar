package internal

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gandarfh/calendar/calendar/internal/calendar"
	"github.com/gandarfh/calendar/calendar/internal/tasks"
	"github.com/gandarfh/calendar/calendar/pkg/constants"
)

type active int

const (
	calendar_page active = iota
	tasks_page
)

type Model struct {
	calendar tea.Model
	tasks    tea.Model
	loaded   bool
	page     active
	keys     KeyMap
}

func New() tea.Model {

	return Model{
		calendar: calendar.New(),
		tasks:    tasks.New(),
		page:     calendar_page,
		loaded:   false,
		keys:     keys,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Calendar):
			return m, tea.Batch(
				m.calendar.(calendar.Model).SetActive(true),
				m.tasks.(tasks.Model).SetActive(false),
			)

		case key.Matches(msg, m.keys.Tasks):
			return m, tea.Batch(
				m.calendar.(calendar.Model).SetActive(false),
				m.tasks.(tasks.Model).SetActive(true),
			)
		}

	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		m.loaded = true
	}

	m.calendar, cmd = m.calendar.Update(msg)
	cmds = append(cmds, cmd)

	m.tasks, cmd = m.tasks.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.loaded {
		return ""
	}

	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.calendar.View(),
		m.tasks.View(),
	)

	return content
}
