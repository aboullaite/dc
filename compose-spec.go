package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type dc struct {
	Version  string
	Networks map[string]network
	Volumes  map[string]volume
	Services map[string]service
}

type network struct {
	Driver, External string
	DriverOpts       map[string]string "driver_opts"
}

type volume struct {
	Driver, External string
	DriverOpts       map[string]string "driver_opts"
}

type service struct {
	ContainerName                     string "container_name"
	Image                             string
	Networks, Ports, Volumes, Command []string
	VolumesFrom                       []string "volumes_from"
	DependsOn                         []string "depends_on"
	CapAdd                            []string "cap_add"
	Build                             struct{ Context, Dockerfile string }
	Environment                       map[string]string
}

func extractComposeSpec(file string) (*dc, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	data := &dc{}
	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil

}
