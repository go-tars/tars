package template

var (
	Main = `package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"{{.Module}}/tars-protocol/{{.App}}"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new({{.Servant}}Imp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("{{.Servant}}Imp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new({{.App}}.{{.Servant}})
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".{{.Servant}}Obj")

	// Run application
	tars.Run()
}
`
)
