package style

import "github.com/charmbracelet/lipgloss"

var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#06D6A0"))

var ChapterStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#9D0006")).
	Background(lipgloss.Color("#FBF1C7"))

var NumberStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#6A7FDB"))

var BookStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#6A7FDB"))

var HeaderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB347")).Bold(true).AlignHorizontal(lipgloss.Center)

var SearchStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB347")).Bold(true)

var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB347")).Align(lipgloss.Center)

var ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6666"))
