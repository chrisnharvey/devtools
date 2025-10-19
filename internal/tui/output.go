package tui

import (
	"io"
	"os"

	"github.com/rivo/tview"
)

type Output struct {
	Command Command
	TApp    *tview.Application
	Values  map[string]any
}

func (o *Output) Render() {
	values := FieldValues{values: o.Values}

	outputView := tview.NewTextView()
	outputView.SetDynamicColors(true).
		SetChangedFunc(func() {
			o.TApp.Draw()
		}).
		SetScrollable(true).
		SetBorder(true).
		SetTitle(" Output ")

	var outputText string

	o.TApp.SetRoot(outputView, true)

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := o.Command.Execute(&values)
	if err != nil {
		outputText = outputText + "[red]Error: " + err.Error() + "\n"
	}

	err = w.Close()
	if err != nil {
		outputText = outputText + "[red]Failed to close writer: " + err.Error() + "\n"
	}
	out, err := io.ReadAll(r)
	if err != nil {
		outputText = outputText + "[red]Failed to read command output: " + err.Error() + "\n"
	}

	os.Stdout = stdout

	outputView.SetText(outputText + "[green]" + string(out))
}
