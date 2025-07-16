package tui

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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

// Defaults
var (
	SearchToggle 		[]string = []string{"tab"}
	SelectDirectory 	[]string = []string{"enter"}
	EnterDirectory		[]string = []string{" ", "/"}
	ExitApplication		[]string = []string{"esc", "q", "ctrl+c"}
	Up					[]string = []string{"up", "k"}
	Down				[]string = []string{"down", "j"}
	Into				[]string = []string{"right", "l", "/", " "}
	OutOf				[]string = []string{"left", "b"}
)

type Keybindings struct {
    SearchToggle     []string `json:"searchToggle"`
    SelectDirectory  []string `json:"selectDirectory"`
    EnterDirectory   []string `json:"enterDirectory"`
    ExitApplication  []string `json:"exitApplication"`
    Up               []string `json:"up"`
    Down             []string `json:"down"`
    Into             []string `json:"into"`
    OutOf            []string `json:"outOf"`
}

func InitialModel() Model {
	err := setKeybindings()
	if err != nil {
		panic(err)
	}

	dirs := GetDirectories(".")
    dir, _ := os.Getwd() // Maybe handle this
	
	ti := textinput.New()
	ti.CharLimit = 50
	ti.Width = 40
	// ti.Focus()

    w, h, _ := term.GetSize(int(os.Stdout.Fd())) // Maybe handle this
	
	return Model{
		directories: dirs,
		allDirectories: dirs,
		selected: make(map[int]string),
		OriginalDir: dir,
		CurrentDir: dir,
		termWidth:  w,
		termHeight: h,
		textInput: ti,
		inputFocused: false, // for whatever reason, this breaks the ui location in the terminal
	}
}

func GetDirectories(rd string) []string {
    entries, err := os.ReadDir(rd)
    if err != nil {
		return []string{}
    }

	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e.Name())
		}
	}

	return dirs
}

func (m Model) Init() tea.Cmd {
	return m.textInput.Cursor.BlinkCmd()
}

func loadKeybindings() (*Keybindings, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return nil, fmt.Errorf("failed to get home directory: %w", err)
    }
    
    configPath := filepath.Join(homeDir, ".config", "gonav", "config.json")
    if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, nil // use defaults if file is not found		
    }
    
    data, err := os.ReadFile(configPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    
	var keybindings Keybindings
    if err := json.Unmarshal(data, &keybindings); err != nil {
        return nil, fmt.Errorf("failed to parse config file: %w", err)
    }

	return &keybindings, nil
}

func setKeybindings() error {
	keybindings, err := loadKeybindings()
	if err != nil {
		return err
	}

	if keybindings == nil {
		return nil
	}

	// use defaults if any are empty
	if len(keybindings.SearchToggle) != 0 {
		SearchToggle = keybindings.SearchToggle
	}

	if len(keybindings.SelectDirectory) != 0 {
		SelectDirectory = keybindings.SelectDirectory
	}

	if len(keybindings.EnterDirectory) != 0 {
		EnterDirectory = keybindings.EnterDirectory
	}

	if len(keybindings.ExitApplication) != 0 {
		ExitApplication = keybindings.ExitApplication
	}

	if len(keybindings.Up) != 0 {
		Up = keybindings.Up
	}

	if len(keybindings.Down) != 0 {
		Down = keybindings.Down
	}

	if len(keybindings.Into) != 0 {
		Into = keybindings.Into
	}

	if len(keybindings.OutOf) != 0 {
		OutOf = keybindings.OutOf
	}
    
    return nil
}
