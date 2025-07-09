package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc": 
			m.CurrentDir = m.OriginalDir
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.directories)-1 {
				m.cursor++
			}

		case "b", "left": 
			nd := GetPreviousDir(m.CurrentDir)	
			m = UpdateModelToNewDir(nd, m.OriginalDir)
		

		case " ", "right":
			if m.cursor < len(m.directories) {
				selectedDir := m.directories[m.cursor]
				dirPath := selectedDir[:len(selectedDir)-1]
				dirPath = fmt.Sprintf("%s/%s", m.CurrentDir, dirPath)
				m = UpdateModelToNewDir(dirPath, m.OriginalDir)
			}

		case "enter": 
			if m.cursor < len(m.directories) {
				selectedDir := m.directories[m.cursor]
				dirPath := selectedDir[:len(selectedDir)-1]
				dirPath = fmt.Sprintf("%s/%s", m.CurrentDir, dirPath)
				m.CurrentDir = dirPath
			}

			tea.ExitAltScreen()
			return m, tea.Quit
		}
	}

	return m, nil
}

func UpdateModelToNewDir(nd string, od string) Model {
	dirs := GetDirectories(nd)

	return Model{
		directories: dirs,
		selected: make(map[int]string),
		OriginalDir: od,
		CurrentDir: nd,
		termWidth:  80,
		termHeight: 24,
	}
}

func GetPreviousDir(dirPath string) string {
	dirs := strings.Split(dirPath, "/")
	dirs = dirs[:len(dirs)-1]
	dirPath = strings.Join(dirs, "/")

	return dirPath
}
