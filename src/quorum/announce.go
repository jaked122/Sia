package quorum

import "common"
import "common/log"
import "network"
import "encoding/json"
import "os"

type Announce struct{
	ID string
	Size int64
	Chunks int16
}

func CreateFileAnnounce(file string) *Announce{
	A:=new(Announce)
	f,err:=os.Open(file)
	if err!=nil{
		panic(err.Error())
	}
	c,err:=f.Stat()
	if err!=nil{
	}
	A.Size=c.Size()
	A.Chunks=15;
	//Chunk File
	return A
}

func (a *Announce) SendOutAnnounce(recipients []*Participant) *Error{
	//TODO Add code to actually send out to participants when it works.
	for i :=range( recipients){
		resp,err:=network.SendMessage(i.Address,i.Port,[]byte(json.Encode(a)))
		if err!=nil{
			panic(err.Error())
		}
		if resp[0]>0{
			//handle file transmission, passing chunk resp[1]
		}
	}
		
	return nil
}















