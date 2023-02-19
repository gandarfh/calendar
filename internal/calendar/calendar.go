package calendar

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gandarfh/calendar/calendar/pkg/constants"
	"github.com/gandarfh/calendar/calendar/pkg/styles"
)

type active bool

type Model struct {
	keys     KeyMap
	help     help.Model
	calendar CalendarControls
	loaded   bool
	active   active
}

func (m Model) SetActive(value bool) tea.Cmd {
	return func() tea.Msg {
		m.active = active(value)
		return active(value)
	}
}

func New() tea.Model {
	return Model{
		active:   true,
		keys:     keys,
		help:     help.New(),
		calendar: NewCalendar(time.Now()),
		loaded:   false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		m.help.Width = msg.Width

		m.loaded = true

	case active:
		m.active = msg

	case CalendarControls:
		m.calendar = msg

	case tea.KeyMsg:
		if m.active {
			switch {
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit

			case key.Matches(msg, m.keys.NextMonth):
				return m, m.calendar.NextMonth()

			case key.Matches(msg, m.keys.PrevMonth):
				return m, m.calendar.PrevMonth()

			case key.Matches(msg, m.keys.NextYear):
				return m, m.calendar.NextYear()

			case key.Matches(msg, m.keys.PrevYear):
				return m, m.calendar.PrevYear()
			}
		}
	}

	return m, nil

}

func (m Model) View() string {
	if !m.loaded {
		return ""
	}

	week := 0
	items := [6][]string{{}, {}, {}, {}, {}}
	for _, v := range m.calendar.Dates() {
		if len(items[week]) == 7 {
			week++
		}

		items[week] = append(items[week], NewDate(v, m.calendar.date).Render())
	}

	dates_view := []string{}
	for _, v := range items {
		dates_view = append(dates_view, lipgloss.JoinHorizontal(lipgloss.Left, v...))
	}

	weeks := []string{
		RenderWeekday("Sunday"),
		RenderWeekday("Monday"),
		RenderWeekday("Tuesday"),
		RenderWeekday("Wednesday"),
		RenderWeekday("Thursday"),
		RenderWeekday("Friday"),
		RenderWeekday("Saturday"),
	}

	calendar := lipgloss.JoinVertical(
		lipgloss.Left,
		m.calendar.Title(),
		"",
		strings.Join(weeks, ""),
		strings.Join(dates_view, "\n"),
	)

	if m.active {
		return styles.Container.Calendar.
			BorderForeground(styles.Theme.Yellow).
			Height(Height()).
			Render(calendar)
	} else {
		return styles.Container.Calendar.
			BorderForeground(styles.Theme.Foreground).
			Height(Height()).
			Render(calendar)
	}

}
