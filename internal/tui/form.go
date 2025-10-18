package tui

import (
	"github.com/chrisnharvey/devtools/pkg/field"
	"github.com/rivo/tview"
)

type Command interface {
	GetFields() []field.Field
	Execute(field.Values) error
}

type Form struct {
	Command Command
	TApp    *tview.Application
	Values  map[string]any
}

func NewForm(command Command, tuiApp *tview.Application) *Form {
	return &Form{
		Command: command,
		TApp:    tuiApp,
		Values:  make(map[string]any),
	}
}

func (f *Form) Render() field.Values {
	form := tview.NewForm()
	for _, fld := range f.Command.GetFields() {
		f.addInputField(form, fld)
	}
	form.AddButton("Execute", func() {
		output := &Output{
			Command: f.Command,
			TApp:    f.TApp,
			Values:  f.Values,
		}

		output.Render()
	})

	f.TApp.SetRoot(form, true)

	return nil
}

func (f *Form) addInputField(form *tview.Form, fld field.Field) {
	switch fld.Type {
	case field.String, field.File:
		form.AddInputField(fld.Description, "", 40, nil, func(text string) {
			f.Values[fld.Name] = text
		})
	case field.TextArea:
		form.AddTextArea(fld.Description, "", 40, 20, 4096, func(text string) {
			f.Values[fld.Name] = text
		})
	}
}
