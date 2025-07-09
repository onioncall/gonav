package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/gonav/tui"
)


func main() {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		// Fallback if /dev/tty unavailable
		return
	}
	defer tty.Close()

	p := tea.NewProgram(
		tui.InitialModel(), 
		tea.WithAltScreen(),
		tea.WithInput(tty),
		tea.WithOutput(tty),
	)

	finalModel, err := p.Run(); 
	if err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}

	m := finalModel.(tui.Model)
	fmt.Print(m.CurrentDir)
}
