package app

import (
	"github.com/spf13/cobra"
)

type App struct {
	comamnds []Command
}

func NewApp() *App {
	return &App{}
}

func (a *App) Add(cmd Command) {
	a.comamnds = append(a.comamnds, cmd)
}

func (a *App) Run() {
	rootCmd := cobra.Command{
		Use:  "devtools",
		RunE: a.runTui,
	}

	// Register all cobra commands
	for _, cmd := range a.comamnds {
		rootCmd.AddCommand(a.createCobraCommand(cmd))
	}

	rootCmd.Execute()
}
