package cli

import "github.com/spf13/pflag"

type FieldValues struct {
	flagSet *pflag.FlagSet
}

func NewFieldValues(flagSet *pflag.FlagSet) *FieldValues {
	return &FieldValues{
		flagSet: flagSet,
	}
}

func (f FieldValues) GetString(name string) string {
	val, _ := f.flagSet.GetString(name)

	return val
}

func (f FieldValues) GetFile(name string) string {
	return f.GetString(name)
}

func (f FieldValues) GetBool(name string) bool {
	val, _ := f.flagSet.GetBool(name)

	return val
}
