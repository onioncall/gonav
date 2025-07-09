package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

type Model struct {
	directories		[]string
	allDirectories	[]string
	cursor   		int
	selected 		map[int]string
	OriginalDir		string
	CurrentDir		string
	termWidth  		int
	termHeight 		int
	textInput    textinput.Model
	Logger		string
	inputFocused bool
}

func InitialModel() Model {
	dirs := GetDirectories(".")
    dir, _ := os.Getwd() // Maybe handle this
	
	ti := textinput.New()
	ti.CharLimit = 50
	ti.Width = 40
	ti.Focus()

    w, h, _ := term.GetSize(int(os.Stdout.Fd())) //Maybe handle this
	
	return Model{
		directories: dirs,
		allDirectories: dirs,
		selected: make(map[int]string),
		OriginalDir: dir,
		CurrentDir: dir,
		termWidth:  w,
		termHeight: h,
		textInput: ti,
		inputFocused: true,
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
	return m.textInput.Cursor.BlinkCmd()
}
