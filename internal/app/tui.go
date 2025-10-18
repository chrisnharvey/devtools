package app

import (
	"github.com/chrisnharvey/devtools/internal/tui"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func (a *App) runTui(c *cobra.Command, args []string) error {
	// Create TUI app??
	tapp := tview.NewApplication()
	menu := tview.NewList()

	// Register all commands
	for _, cmd := range a.comamnds {
		menu.AddItem(cmd.GetName(), cmd.GetDescription(), 0, func() {
			form := tui.NewForm(cmd, tapp)

			form.Render()
		})
	}

	tapp.SetRoot(menu, true)

	return tapp.Run()
	//return nil
}
