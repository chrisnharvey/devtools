package main

import (
	"github.com/chrisnharvey/devtools/cmd/charcount"
	"github.com/chrisnharvey/devtools/cmd/jsonminify"
	"github.com/chrisnharvey/devtools/cmd/jsonprettify"
	"github.com/chrisnharvey/devtools/cmd/sha256"
	"github.com/chrisnharvey/devtools/cmd/urldecode"
	"github.com/chrisnharvey/devtools/cmd/urlencode"
	"github.com/chrisnharvey/devtools/internal/app"
)

func main() {
	a := app.NewApp()

	a.Add(charcount.New())
	a.Add(sha256.New())
	a.Add(urlencode.New())
	a.Add(urldecode.New())
	a.Add(jsonminify.New())
	a.Add(jsonprettify.New())

	a.Run()
}
