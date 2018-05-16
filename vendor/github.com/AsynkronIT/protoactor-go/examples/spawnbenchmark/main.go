package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/mailbox"
)

type request struct {
	num  int
	size int
	div  int
}

var (
	props = actor.FromProducer(newState).WithMailbox(mailbox.UnboundedLockfree())
)

type state struct {
	sum     int
	replies int
	replyTo *actor.PID
}

func newState() actor.Actor {
	return &state{}
}

func (s *state) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *request:
		if msg.size == 1 {
			ctx.Respond(msg.num)
			return
		}

		s.replies = msg.div
		s.replyTo = ctx.Sender()
		for i := 0; i < msg.div; i++ {
			child := actor.Spawn(props)
			child.Request(&request{
				num:  msg.num + i*(msg.size/msg.div),
				size: msg.size / msg.div,
				div:  msg.div,
			}, ctx.Self())
		}
	case int:
		s.sum += msg
		s.replies--
		if s.replies == 0 {
			s.replyTo.Tell(s.sum)
		}
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write mem profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GC()

	start := time.Now()
	pid := actor.Spawn(props)
	res, _ := pid.RequestFuture(&request{
		num:  0,
		size: 1000000,
		div:  10,
	}, 10*time.Second).Result()
	result := res.(int)

	took := time.Since(start)
	fmt.Printf("Result: %d in %d ms.\n", result, took.Nanoseconds()/1e6)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
