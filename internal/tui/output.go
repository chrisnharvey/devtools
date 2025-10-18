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

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	o.Command.Execute(&values)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = stdout

	outputView := tview.NewTextView()
	outputView.SetDynamicColors(true).
		SetText(string(out)).
		SetChangedFunc(func() {
			o.TApp.Draw()
		}).
		SetScrollable(true).
		SetBorder(true).
		SetTitle(" Output ")

	o.TApp.SetRoot(outputView, true)
}
