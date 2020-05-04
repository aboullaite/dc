package main

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

// Docker compose executable
const DCExec = "docker-compose"

func execCmdHost(arg ...string) {
	cmd := exec.Command(DCExec, arg...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("could not get stderr pipe: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("could not get stdout pipe: %v", err)
	}
	go func() {
		merged := io.MultiReader(stderr, stdout)
		scanner := bufio.NewScanner(merged)
		for scanner.Scan() {
			msg := scanner.Text()
			log.Printf("%s", msg)
		}
	}()
	if err := cmd.Start(); err != nil {
		log.Fatalf("could not run cmd: %v", err)
	}
	if err != nil {
		log.Fatalf("could not wait for cmd: %v", err)
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

func logs(file, service string) {
	execCmdHost("-f", file, "logs", "-f", service)
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

func checkDCExists() {
	_, err := exec.LookPath(DCExec)
	if err != nil {
		log.Fatalln("Can't find docker compose executable! make sure it's installed")
	}
}
