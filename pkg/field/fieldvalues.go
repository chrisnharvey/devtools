package field

type Values interface {
	GetString(name string) string
	GetFile(name string) string
	GetBool(name string) bool
}
