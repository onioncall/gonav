package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	directories		[]string
	cursor   		int
	selected 		map[int]string
	OriginalDir		string
	CurrentDir		string
	termWidth  		int
	termHeight 		int
	textInput    textinput.Model
	inputFocused bool
}

func InitialModel() Model {
	dirs := GetDirectories(".")
    dir, _ := os.Getwd() // maybe handle this

	return Model{
		directories: dirs,
		selected: make(map[int]string),
		OriginalDir: dir,
		CurrentDir: dir,
		termWidth:  80,
		termHeight: 24,
	}
}

func GetDirectories(rd string) []string {
    entries, err := os.ReadDir(rd)
    if err != nil {
		s := make([]string, 0, 0)
		return s
    }

	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dir := fmt.Sprintf("%s%s", e.Name(), "/")
			dirs = append(dirs, dir)
		}
	}

	return dirs
}

func (m Model) Init() tea.Cmd {
	return nil
}
