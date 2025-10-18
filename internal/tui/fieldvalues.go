package tui

type FieldValues struct {
	values map[string]any
}

func (f *FieldValues) GetString(name string) string {
	val, ok := f.values[name]
	if !ok {
		return ""
	}
	return val.(string)
}

func (f *FieldValues) GetFile(name string) string {
	return f.GetString(name)
}

func (f *FieldValues) GetBool(name string) bool {
	val, ok := f.values[name]
	if !ok {
		return false
	}
	return val.(bool)
}
