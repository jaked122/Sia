package main

import (
	"network"
	"quorum"
	"time"
)

func main() {
	// ms == messageSender
	ms, err := network.NewTCPServer(9988)
	if err != nil {
		println("fail")
	}

	// mh == messageHandler
	mh0, err := quorum.CreateState(ms, 0)
	if err != nil {
		println("fail")
	}

	mh1, err := quorum.CreateState(ms, 1)
	if err != nil {
		println("fail")
	}

	mh0.AddParticipant(mh1.Self(), 1)
	mh1.AddParticipant(mh0.Self(), 0)

	mh0.Start()
	mh1.Start()

	time.Sleep(time.Second)
}
