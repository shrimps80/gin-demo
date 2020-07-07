package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func init() {
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		fmt.Println("every 10 seconds executing")
	})
	
	go c.Start()
	defer c.Stop()
}
