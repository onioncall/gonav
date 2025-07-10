package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.inputFocused {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "tab":
				m.inputFocused = false
				return m, nil
			case " ":
				if len(m.directories) == 1 {
					dirPath := fmt.Sprintf("%s/%s", m.CurrentDir, m.directories[0])
					m = m.UpdateModelToNewDir(dirPath)
					m.inputFocused = true
					m.textInput.Focus()
					return m, nil
				}
			case "enter":
				if len(m.directories) == 1 {
					dirPath := fmt.Sprintf("%s/%s", m.CurrentDir, m.directories[0])
					m.CurrentDir = dirPath
					tea.ExitAltScreen()
					return m, tea.Quit
				}					
				return m, nil
			}
		}

		m.textInput, cmd = m.textInput.Update(msg)
		m.FilterDirectories()
		return m, cmd
	}
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc": 
			m.CurrentDir = m.OriginalDir
			return m, tea.Quit
		case "s", "tab":
			m.inputFocused = true
			m.textInput.Focus()
			return m, m.textInput.Cursor.BlinkCmd()
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.directories)-1 {
				m.cursor++
			}
		case "b", "left": 
			lastDirInPath := filepath.Base(m.CurrentDir)
			nd := GetPreviousDir(m.CurrentDir)	
			m = m.UpdateModelToNewDir(nd)
			m.cursor = m.GetCursorReturnPosition(lastDirInPath)
		case " ", "right":
			if m.cursor < len(m.directories) {
				selectedDir := m.directories[m.cursor]
				dirPath := fmt.Sprintf("%s/%s", m.CurrentDir, selectedDir)
				m = m.UpdateModelToNewDir(dirPath)
			}
		case "enter": 
			if m.cursor < len(m.directories) {
				selectedDir := m.directories[m.cursor]
				dirPath := fmt.Sprintf("%s/%s", m.CurrentDir, selectedDir)
				m.CurrentDir = dirPath
			}

			tea.ExitAltScreen()
			return m, tea.Quit
		}
    case tea.WindowSizeMsg:
        m.termWidth = msg.Width
        m.termHeight = msg.Height
        return m, nil
    }
	return m, nil
}

func (m *Model) FilterDirectories() {
	searchTerm := strings.ToLower(m.textInput.Value())
	
	if searchTerm != "" {
		m.directories = make([]string, 0)
		for _, dir := range m.allDirectories {
			if strings.Contains(strings.ToLower(dir), searchTerm) {
				m.directories = append(m.directories, dir)
			}
		}
	} else {
		m.directories = m.allDirectories
	}


	if m.cursor >= len(m.directories) {
		m.cursor = 0
		if len(m.directories) > 0 {
			m.cursor = len(m.directories) - 1
		}
	}
}

func (m Model) UpdateModelToNewDir(nd string) Model {
	dirs := GetDirectories(nd)
	ti := textinput.New()
	ti.CharLimit = 50
	ti.Width = 40
	
	return Model{
		directories: dirs,
		allDirectories: dirs,
		selected: make(map[int]string),
		OriginalDir: m.OriginalDir,
		CurrentDir: nd,
		termWidth: m.termWidth,
		termHeight: m.termHeight,
		textInput: ti,
		inputFocused: false,
	}
}

func GetPreviousDir(dirPath string) string {
	dirs := strings.Split(dirPath, "/")
	dirs = dirs[:len(dirs)-1]
	dirPath = strings.Join(dirs, "/")

	return dirPath
}

func (m Model) GetCursorReturnPosition(lastDirInPath string) int {
	for i, dir := range m.directories {
		if dir == lastDirInPath {
			return i
		} 
	}

	return 0
}
