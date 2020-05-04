package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

// Docker compose executable
const DCExec = "docker-compose"

func execCmdHost(arg ...string) {
	cmd := exec.Command(DCExec, arg...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("could not get stdout pipe: %v", err)
	}
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Println(msg)
		}
	}()
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error Running cmd: %v", err)
	}
}

func up(file string) {
	execCmdHost("-f", file, "up")
}

func upBuild(file, service string) {
	execCmdHost("-f", file, "up", "--build", service)
}

func stop(file, service string) {
	execCmdHost("-f", file, "stop", service)
}

func down(file string) {
	execCmdHost("-f", file, "down")
}

func logs(file string) {
	execCmdHost("-f", file, "logs", "-f")
}

func build(file, service string) {
	execCmdHost("-f", file, "build", service)
}

func start(file, service string) {
	execCmdHost("-f", file, "start", service)
}

func refresh(file, service string) {
	stop(file, service)
	upBuild(file, service)
}

func bootstrap(file string) {
	up(file)
}

func checkDCExists() {
	_, err := exec.LookPath(DCExec)
	if err != nil {
		log.Fatalln("Can't find docker compose executable! make sure it's installed")
	}
}
