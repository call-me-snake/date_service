package main

import (
	"fmt"
	"github.com/call-me-snake/date_service/internal/httpServer"
	"github.com/call-me-snake/date_service/internal/model"
	"github.com/jessevdk/go-flags"
	"log"
)

type envs struct {
	ServerPort int `long:"port" env:"PORT" description:"port of microservice" default:"8002"`
}

func getEnvs() (envs, error) {
	e := envs{}
	var err error
	parser := flags.NewParser(&e, flags.Default)
	if _, err = parser.Parse(); err != nil {
		return e, fmt.Errorf("Init: %v", err)
	}
	return e, nil
}

func main() {
	log.Print("Started")
	environments, err := getEnvs()
	if err != nil {
		log.Fatal(err)
	}
	s := httpServer.New(fmt.Sprintf(":%d", environments.ServerPort))
	err = s.Start(&model.TimeData{})
	if err != nil {
		log.Fatal(err)
	}
}
