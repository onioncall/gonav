package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var content string
	content += fmt.Sprintf("%s\n\n", m.CurrentDir)
	if len(m.directories) == 0 {
		content += fmt.Sprintf("Empty Directory")
	} else {
		for i, choice := range m.directories {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			content += fmt.Sprintf("%s - %s\n", cursor, choice)
		}
	}

	box := lipgloss.NewStyle().
		Height(24).
		Width(84).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2).
		Render(content)

	inputStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(0, 1).
		Width(50)

	inputContent := fmt.Sprintf("Search: %s", m.textInput.View())
	inputBox := inputStyle.Render(inputContent)

	combined := lipgloss.JoinVertical(
		lipgloss.Center,
		box,
		"",
		inputBox,
	)

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center,
		lipgloss.Center,
		combined,
	)
}
