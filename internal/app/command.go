package app

import "github.com/chrisnharvey/devtools/pkg/field"

type Command interface {
	GetUse() string
	GetDescription() string
	GetFields() []field.Field
	Execute(field.Values) error
}
