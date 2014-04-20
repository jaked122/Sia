package main

import (
	"network"
	"quorum"
	"time"
)

func main() {
	println("starting")

	// create a networking server that can pass messages between hosts
	tcp, err := network.NewTCPServer(9988)
	if err != nil {
		println(err)
		return
	}

	// create states that communicate over the tcp server
	s0, _ := quorum.CreateState(tcp, 0)
	s1, _ := quorum.CreateState(tcp, 1)

	// add states to networking server
	tcp.MessageHandlers[0] = &s0
	tcp.MessageHandlers[1] = &s1

	// add states to each other
	s0.AddParticipant(s1.Self(), 1)
	s1.AddParticipant(s0.Self(), 0)

	// start the algorithms for each state
	s0.Start()
	s1.Start()

	time.Sleep(time.Second)

	println("done")
}
