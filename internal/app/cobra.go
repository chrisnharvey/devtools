package app

import (
	"github.com/chrisnharvey/devtools/internal/cli"
	"github.com/chrisnharvey/devtools/pkg/field"
	"github.com/spf13/cobra"
)

func (a *App) createCobraCommand(cmd Command) *cobra.Command {
	cobraCmd := &cobra.Command{
		Use:   cmd.GetUse(),
		Short: cmd.GetDescription(),
		RunE: func(c *cobra.Command, args []string) error {
			values := cli.NewFieldValues(c.Flags())
			return cmd.Execute(values)
		},
	}

	for _, f := range cmd.GetFields() {
		switch f.Type {
		case field.String, field.TextArea, field.File:
			cobraCmd.Flags().String(f.Name, "", f.Description)
			if f.Required {
				_ = cobraCmd.MarkFlagRequired(f.Name)
			}
		case field.Bool:
			cobraCmd.Flags().Bool(f.Name, false, f.Description)
		}
	}

	return cobraCmd
}
