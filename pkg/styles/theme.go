package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type theme struct {
	Background lipgloss.AdaptiveColor
	Foreground lipgloss.AdaptiveColor
	Red        lipgloss.AdaptiveColor
	Green      lipgloss.AdaptiveColor
	Yellow     lipgloss.AdaptiveColor
	Blue       lipgloss.AdaptiveColor
	Gray       lipgloss.AdaptiveColor
	DarkGray   lipgloss.AdaptiveColor
}

var Theme *theme = &theme{
	Background: lipgloss.AdaptiveColor{Light: "000", Dark: "000"},
	Foreground: lipgloss.AdaptiveColor{Light: "015", Dark: "015"},
	Red:        lipgloss.AdaptiveColor{Light: "001", Dark: "001"},
	Green:      lipgloss.AdaptiveColor{Light: "002", Dark: "002"},
	Yellow:     lipgloss.AdaptiveColor{Light: "003", Dark: "003"},
	Blue:       lipgloss.AdaptiveColor{Light: "004", Dark: "004"},
	Gray:       lipgloss.AdaptiveColor{Light: "244", Dark: "244"},
	DarkGray:   lipgloss.AdaptiveColor{Light: "234", Dark: "234"},
}
