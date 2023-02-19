package calendar

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gandarfh/calendar/calendar/pkg/constants"
	"github.com/gandarfh/calendar/calendar/pkg/styles"
	"github.com/gandarfh/calendar/calendar/pkg/utils"
)

func Height() int {
	return (constants.WindowSize.Height - (constants.WindowSize.Height / 6)) / 6
}

func Width() int {
	return (constants.WindowSize.Width - (constants.WindowSize.Width / 4)) / 7
}

type CalendarControls struct {
	date time.Time
}

func NewCalendar(date time.Time) CalendarControls {
	return CalendarControls{date}
}

func (c CalendarControls) Date() time.Time {
	return c.date
}

func (c CalendarControls) Title() string {
	content := lipgloss.NewStyle().
		Bold(true).
		Foreground(styles.Theme.Yellow)

	return content.Render(c.date.Format("January - 2006"))
}

func (c CalendarControls) NextMonth() tea.Cmd {
	return func() tea.Msg {
		c.date = c.date.AddDate(0, 1, 0)
		return c
	}
}
func (c CalendarControls) PrevMonth() tea.Cmd {
	return func() tea.Msg {
		c.date = c.date.AddDate(0, -1, 0)
		return c
	}
}

func (c CalendarControls) NextYear() tea.Cmd {
	return func() tea.Msg {
		c.date = c.date.AddDate(1, 0, 0)
		return c
	}
}
func (c CalendarControls) PrevYear() tea.Cmd {
	return func() tea.Msg {
		c.date = c.date.AddDate(-1, 0, 0)
		return c
	}
}

func (c CalendarControls) Dates() []time.Time {
	now := utils.BeginningOfMonth(c.date).AddDate(0, 0, 0)

	start := utils.BeginningOfWeek(utils.BeginningOfMonth(now))
	end := utils.EndOfWeek(utils.EndOfMonth(now)).AddDate(0, 1, 0)

	dates := utils.RangeBetweenDates(start, end)

	return dates[:42]
}

type Date struct {
	date   time.Time
	parent time.Time
}

func NewDate(date, parent time.Time) Date {
	return Date{date, parent}
}

func (d Date) Events() []string {
	events_loaded := []string{"Dormir", "Comer", "Seila"}

	events := []string{}

	for _, t := range events_loaded {
		item := lipgloss.NewStyle().Foreground(styles.Theme.Blue)
		if d.isSunday() || d.isSaturday() || !d.isSameMonth() {
			item.Foreground(styles.Theme.DarkGray)
		}
		events = append(events, item.Render(t))
	}

	return events
}

func (d Date) Tasks() []string {
	tasks_loaded := []string{"Test"}

	tasks := []string{}

	for _, t := range tasks_loaded {
		item := lipgloss.NewStyle().Foreground(styles.Theme.Green)
		if d.isSunday() || d.isSaturday() || !d.isSameMonth() {
			item.Foreground(styles.Theme.DarkGray)
		}

		tasks = append(tasks, item.Render(t))
	}

	return tasks
}

func (d Date) List() []string {
	items := []string{}

	items = append(items, d.Tasks()...)
	items = append(items, d.Events()...)

	if len(items) >= Height()-2 {
		items[Height()-3] = lipgloss.NewStyle().Foreground(styles.Theme.DarkGray).Render("...")

		return items[:Height()-2]
	}

	return items
}

func (d Date) isSameMonth() bool {
	return d.parent.Month() == d.date.Month()
}

func (d Date) isSunday() bool {
	return int(d.date.Weekday()) == 0
}

func (d Date) isSaturday() bool {
	return int(d.date.Weekday()) == 6
}

func (d Date) isToday() bool {
	return d.date.Format("2006-01-02") == time.Now().Format("2006-01-02")
}

func (d Date) Render() string {

	label := lipgloss.NewStyle().
		Foreground(styles.Theme.Foreground)

	if d.isSunday() || d.isSaturday() {
		label.Foreground(styles.Theme.Red)
	}

	if d.isToday() {
		label.Bold(true).Foreground(styles.Theme.Green)

		items := []string{label.Render(d.date.Format("02") + utils.TruncateSimple(" - Today", Width()-3))}
		items = append(items, d.List()...)

		return lipgloss.NewStyle().
			Height(Height()).
			Width(Width()).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Left,
					items...,
				),
			)
	}

	if !d.isSameMonth() {
		label.Foreground(styles.Theme.DarkGray)

		items := []string{label.Render(d.date.Format("02"))}
		items = append(items, d.List()...)

		return lipgloss.NewStyle().
			Height(Height()).
			Width(Width()).
			Render(
				lipgloss.JoinVertical(
					lipgloss.Left,
					items...,
				),
			)
	}

	items := []string{label.Render(d.date.Format("02"))}
	items = append(items, d.List()...)

	return lipgloss.NewStyle().
		Height(Height()).
		Width(Width()).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				items...,
			),
		)

}

func RenderWeekday(weekday string) string {

	isSunday := weekday == "Sunday"
	isSaturday := weekday == "Saturday"
	content := lipgloss.NewStyle().
		Width(Width()).
		Bold(true).
		Foreground(styles.Theme.Gray)

	if isSunday || isSaturday {
		content.Foreground(styles.Theme.Red)
	}

	return content.Render(utils.TruncateSimple(weekday, Width()-2))
}
