package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rjeczalik/notify"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dc",
		Usage: "Continuous deployment for your Docker compose services during development",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   "docker-compose.yaml",
				Usage:   "Docker compose file",
			},
		},
		Action: func(c *cli.Context) error {
			checkDCExists()
			composefile := c.String("file")
			dc, err := extractComposeSpec(composefile)
			up(composefile)
			if err != nil {
				return err
			}

			for k, v := range dc.Services {
				path, _ := filepath.Abs(v.Build.Context)
				for {
					watchFolder(path, composefile, k)
				}
			}
			defer down(composefile)
			return nil

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func watchFolder(projectPath, composeFile, service string) {
	channel := make(chan notify.EventInfo, 1)
	if err := notify.Watch(projectPath, channel, notify.All); err != nil {
		log.Fatalln(err)
	}
	defer notify.Stop(channel)
	<-channel
	//log.Printf("File %s changed", ei.Event().String())
	refresh(filepath.Join(projectPath, composeFile), service)
}
