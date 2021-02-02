package test3

import (
	"fmt"
	"time"
)

type (
	Cronjob interface {
		GetName()
	}

	cronjob struct {
		name string
		latestExecuteTime time.Time
	}
)

func (c cronjob) GetName() {
	fmt.Println("hello, my name is ", c.name, "cronJob")
}

func NewCronjob(name string, latestExecuteTime time.Time) Cronjob {
	return &cronjob{
		name: name,
		latestExecuteTime: latestExecuteTime,
	}
}
