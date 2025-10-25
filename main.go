package main

import (
	"github.com/chrisnharvey/devtools/cmd/base64decode"
	"github.com/chrisnharvey/devtools/cmd/base64encode"
	"github.com/chrisnharvey/devtools/cmd/charcount"
	"github.com/chrisnharvey/devtools/cmd/hmac"
	"github.com/chrisnharvey/devtools/cmd/htmldecode"
	"github.com/chrisnharvey/devtools/cmd/htmlencode"
	"github.com/chrisnharvey/devtools/cmd/jsonminify"
	"github.com/chrisnharvey/devtools/cmd/jsonprettify"
	"github.com/chrisnharvey/devtools/cmd/jsontoyaml"
	"github.com/chrisnharvey/devtools/cmd/jwt"
	"github.com/chrisnharvey/devtools/cmd/sha256"
	"github.com/chrisnharvey/devtools/cmd/timestamp"
	"github.com/chrisnharvey/devtools/cmd/ulid"
	"github.com/chrisnharvey/devtools/cmd/urldecode"
	"github.com/chrisnharvey/devtools/cmd/urlencode"
	"github.com/chrisnharvey/devtools/cmd/uuid"
	"github.com/chrisnharvey/devtools/cmd/yamltojson"
	"github.com/chrisnharvey/devtools/internal/app"
)

func main() {
	a := app.NewApp()

	a.Add(base64decode.New())
	a.Add(base64encode.New())
	a.Add(charcount.New())
	a.Add(hmac.New())
	a.Add(htmldecode.New())
	a.Add(htmlencode.New())
	a.Add(jsontoyaml.New())
	a.Add(jwt.New())
	a.Add(jsonminify.New())
	a.Add(jsonprettify.New())
	a.Add(sha256.New())
	a.Add(timestamp.New())
	a.Add(ulid.New())
	a.Add(urlencode.New())
	a.Add(urldecode.New())
	a.Add(uuid.New())
	a.Add(yamltojson.New())

	a.Run()
}
