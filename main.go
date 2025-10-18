package main

import (
	"github.com/chrisnharvey/devtools/cmd/sha256"
	"github.com/chrisnharvey/devtools/internal/app"
)

func main() {
	a := app.NewApp()

	a.Add(sha256.New())

	a.Run()
}
