package styles

import "github.com/charmbracelet/lipgloss"

type container struct {
	Base     lipgloss.Style
	Calendar lipgloss.Style
	Tasks    lipgloss.Style
}

var Container = func() container {
	base := lipgloss.NewStyle().Margin(1)
	Calendar := lipgloss.NewStyle().Margin(1).Padding(0, 1).
		Border(lipgloss.RoundedBorder())

	Tasks := lipgloss.NewStyle().Margin(1, 0).Padding(0, 1).
		Border(lipgloss.RoundedBorder())

	return container{
		Base:     base,
		Calendar: Calendar,
		Tasks:    Tasks,
	}
}()
