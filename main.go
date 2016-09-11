package main

//go:generate statik -src=.\static
//go generate go get github.com/jteeuwen/go-bindata/...
//go:generate go-bindata -o settings.go -pkg main settings/...

import (
	"encoding/json"
	"os"

	"github.com/enlivengo/admin"
	"github.com/enlivengo/database"
	"github.com/enlivengo/enliven"
	"github.com/enlivengo/enliven/config"
	"github.com/enlivengo/enliven/middleware/session"
	_ "github.com/enlivengo/skeleton/statik"
	"github.com/enlivengo/static"
	"github.com/enlivengo/statik"
	"github.com/enlivengo/user"
)

func getConfig() config.Config {
	environment := "development"
	args := os.Args[1:]
	if len(args) > 0 {
		environment = args[0]
	}
	settingsBytes := MustAsset("settings/" + environment + ".json")

	var enlivenConfig config.Config

	if err := json.Unmarshal(settingsBytes, &enlivenConfig); err != nil {
		panic(err)
	}

	return enlivenConfig
}

func main() {
	ev := enliven.New(getConfig())

	ev.AddMiddleware(session.NewMemoryStorageMiddleware())
	ev.AddApp(static.NewApp())
	ev.AddApp(statik.NewApp())
	ev.AddApp(database.NewApp())
	ev.AddApp(user.NewApp())
	ev.AddApp(admin.NewApp())

	ev.AddRoute("/", rootView)

	ev.Run()
}
