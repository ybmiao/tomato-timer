package main

import (
	"countdown/internal/cmd"
	"countdown/internal/conf"
	"github.com/urfave/cli"
	"os"
	log "unknwon.dev/clog/v2"
)

func init() {
	conf.App.Name = "Tomato timing"
	conf.App.Usage = "A tomato timing tool"
	conf.App.Version = "1.0.0"
}

func main() {
	app := cli.NewApp()
	app.Name = conf.App.Name
	app.Usage = conf.App.Usage
	app.Version = conf.App.Version
	app.Action = cmd.TomatoTimer

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}

}
