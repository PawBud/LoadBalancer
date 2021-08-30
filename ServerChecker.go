package main

import (
	"github.com/go-co-op/gocron"
	"time"
)

func startServerCheck() {
	s:= gocron.NewScheduler(time.Local)
	for _, server := range serverList {
		s.Every(2).Second().Do(server.serverStatus)
	}
}