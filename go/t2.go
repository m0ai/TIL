package main

import (
	"fmt"
	"time"
)

type (
	Cronjob struct {
		Name string
		LatestExecuteTime time.Time
	}
)

func (c Cronjob) Run() {
	fmt.Println("Hello, I'm Cronjob", c.Name)
}

func NewCronjob(name string, latestExecuteTime time.Time) *Cronjob {
	return &Cronjob{Name: "KoreanWeatherSource", LatestExecuteTime: nil}
}



