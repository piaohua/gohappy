package plugin

import (
	"log"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type PassivationAware interface {
	Init(*actor.PID, time.Duration)
	Reset(time.Duration)
	Cancel()
}

type PassivationHolder struct {
	timer *time.Timer
	done  bool
}

func (state *PassivationHolder) Reset(duration time.Duration) {
	if state.timer == nil {
		log.Fatalf("Cannot reset passivation of a non-started actor")
	}
	if !state.done {
		state.timer.Reset(duration)
	}
}

func (state *PassivationHolder) Init(pid *actor.PID, duration time.Duration) {
	state.timer = time.NewTimer(duration)
	state.done = false
	go func() {
		select {
		case <-state.timer.C:
			pid.Stop()
			state.done = true
			break
		}
	}()
}

func (state *PassivationHolder) Cancel() {
	if state.timer != nil {
		state.timer.Stop()
	}
}

type PassivationPlugin struct {
	Duration time.Duration
}

func (pp *PassivationPlugin) OnStart(ctx actor.Context) {
	if a, ok := ctx.Actor().(PassivationAware); ok {
		a.Init(ctx.Self(), pp.Duration)
	}
}

func (pp *PassivationPlugin) OnOtherMessage(ctx actor.Context, msg interface{}) {
	if p, ok := ctx.Actor().(PassivationAware); ok {
		switch msg.(type) {
		case *actor.Stopped:
			p.Cancel()
		default:
			p.Reset(pp.Duration)
		}
	}
}
