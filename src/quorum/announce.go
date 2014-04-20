package quorum

import (
	"common"
	"common/erasure"
	"common/log"
	"encoding/json"
	"fmt"
	"io"
	"network"
	"os"
)

type File struct {
	ID                string
	Size              int
	Chunkdistribution map[int]*[]Participant
}

type Announce struct {
	ID          string
	Size        int
	Chunks      int
	networkinfo File
}

func CreateFileAnnounce(file string) *Announce {
	A := new(Announce)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	c, err := f.Stat()
	if err != nil {
	}
	A.Size = int(c.Size())
	//this should be enough.
	A.Chunks = A.Size / (1 << 20)
	//Chunk File
	d := make([]byte, A.Size/(1<<20))
	indexs := 0
	//generate ID and dump the chunks into files
	//A.ID=""
	for _ = 0; err != io.EOF; _, err = f.Read(d) {
		if err != nil {
		}
		redunt, _ := erasure.EncodeRing(A.Chunks, A.Size/A.Chunks, d)
		A.Chunks = len(redunt)
		indexs = indexs + 1
		for index, val := range redunt {
			k, e := os.Open(fmt.Sprintf("%s-%d%d", A.ID, indexs, index))
			if e != nil {
				log.Fatal(e)
			}
			if _, e = k.Write([]byte(val)); e != nil {
				log.Fatal(e)
			}
		}
	}

	return A
}
func (a *Announce) Identifier() byte {
	return 2
}
func (a *Announce) HandleMessage(c []byte) {
	//do stuff
}

func (a *Announce) Type() string {
	return "Announce"
}

func (a *Announce) SendOutAnnounce(recipients []*Participant) error {
	//TODO Add code to actually send out to participants when it works.
	server, err := network.NewTCPServer(7777)
	if err != nil {
		log.Fatal("TCP Server not initialized")
	}
	for _, i := range recipients {
		s, err := json.Marshal(a)
		s = []byte(string(rune(0)) + string(s))
		if err != nil {
			log.Fatal(err)
		}
		c := new(common.Message)
		c.Payload = s
		c.Destination = i.Address
		err = server.SendMessage(c)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
