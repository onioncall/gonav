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
		panic(err)
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

	// if m.Logger != "" {
	// 	fmt.Println(m.Logger)
	// }

	fmt.Print(m.CurrentDir)
}
