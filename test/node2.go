package main

import (
	"log"
	"time"

	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/examples/chat/messages"
	"github.com/AsynkronIT/protoactor-go/remote"
)

func main() {
	remote.Start("127.0.0.1:0")

	server := actor.NewPID("127.0.0.1:8080", "chatserver")
	//server1 := actor.NewPID("127.0.0.1:8080", "chatserver")
	//server2 := actor.NewPID("127.0.0.1:8080", "chatserver")
	//spawn our chat client inline
	props := actor.FromFunc(func(context actor.Context) {
		switch msg := context.Message().(type) {
		case *messages.Connected:
			log.Println(msg.Message)
			context.SetReceiveTimeout(10 * time.Second)
		case *actor.ReceiveTimeout:
			log.Printf("ReceiveTimeout: %v", context.Self().String())
			context.Self().Stop()
		case *messages.SayResponse:
			log.Printf("%v: %v", msg.UserName, msg.Message)
		case *messages.NickResponse:
			log.Printf("%v is now known as %v", msg.OldUserName, msg.NewUserName)
		default:
			log.Printf("unknown message %#v\n", msg)
		}
	})

	client := actor.Spawn(props)

	server.Tell(&messages.Connect{
		Sender: client,
	})
	log.Println("stoping")
	//client.Stop()
	log.Println("stoped")

	//client1 := actor.Spawn(props)

	//server1.Tell(&messages.Connect{
	//	Sender: client1,
	//})

	//client2 := actor.Spawn(props)

	//server2.Tell(&messages.Connect{
	//	Sender: client2,
	//})

	nick := "Roger"
	cons := console.NewConsole(func(text string) {
		server.Tell(&messages.SayRequest{
			UserName: nick,
			Message:  text,
		})
	})
	//write /nick NAME to change your chat username
	cons.Command("/nick", func(newNick string) {
		server.Tell(&messages.NickRequest{
			OldUserName: nick,
			NewUserName: newNick,
		})
		nick = newNick
	})
	cons.Run()
}
