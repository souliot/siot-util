package sutil

import (
	"github.com/robfig/cron/v3"
)

func init() {

}

var (
	cs = make(chan *Cron)
	c  = cron.New()
)

type Cron struct {
	spec string
	fn   func()
}

func newCron(spec string, fn func()) *Cron {
	return &Cron{
		spec: spec,
		fn:   fn,
	}
}

func (m *Cron) Do() {
	c.AddFunc(m.spec, m.fn)
}

func StartCrons() {
	go start()
}

func start() {
	c.Start()
	for {
		select {
		case c := <-cs:
			c.Do()
		}
	}
}

func StopCrons() {
	c.Stop()
}

func ResetCrons() {
	c.Stop()
	c = cron.New()
}

func AddCron(spec string, fn func()) {
	cs <- newCron(spec, fn)
}
