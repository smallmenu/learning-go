package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/x-funs/go-fun"
)

var (
	c = cron.New()
)

type Delay struct {
	Name  string
	Count int
}

func (p *Delay) Run() {
	p.Count++
	//if p.Count == 2 {
	//	panic("panic")
	//}
	fmt.Println(fun.Date(fun.DatetimeMilliPattern), p.Name, p.Count)
	randomInt := fun.RandomInt64(4, 10)
	time.Sleep(time.Second * time.Duration(randomInt))
	if p.Count >= 10 {
		fmt.Println("remove and stop cron")
		entries := c.Entries()
		for _, entry := range entries {
			c.Remove(entry.ID)
		}
		c.Stop()
	}
}

func main() {
	c.AddJob("@every 3s", cron.NewChain(cron.Recover(cron.DefaultLogger), cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&Delay{Name: "DelayJob"}))

	c.Start()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
