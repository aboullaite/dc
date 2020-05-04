package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

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
			setupStopHandler(composefile)
			dc, err := extractComposeSpec(composefile)
			bootstrap(composefile)
			if err != nil {
				return err
			}
			done := make(chan bool)
			for k, v := range dc.Services {
				if v.Build.Context != "" {
					path, _ := filepath.Abs(v.Build.Context)
					go WatchFolder(path, composefile, k)
				}
			}
			<-done

			return nil

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// WatchFolder track changes in a particular service folder, this is usually the cotext path
// and triggers build and deployement
func WatchFolder(projectPath, composeFile, service string) {
	for {
		channel := make(chan notify.EventInfo, 1)
		if err := notify.Watch(projectPath, channel, notify.All); err != nil {
			log.Fatalln(err)
		}
		defer notify.Stop(channel)
		<-channel
		log.Printf("changes detected in service %s, redeploying ...", service)
		refresh(composeFile, service)
	}
}

// SetupStopHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS.
func setupStopHandler(composeFile string) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		down(composeFile)
		os.Exit(0)
	}()
}
