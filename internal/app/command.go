package app

import "github.com/chrisnharvey/devtools/pkg/field"

type Command interface {
	GetName() string
	GetDescription() string
	GetUse() string
	GetFields() []field.Field
	Execute(field.Values) error
}
