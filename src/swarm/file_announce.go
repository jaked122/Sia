package swarm
import (
	"common"
	"encoding/json"
)
const(
	//Request a file from the network
	RequestFile = iota
	//Serve a file for the network to obtain
	ServeFile
	//Request a particular chunk to be stored
	RequestChunk
)
type FileAnnounce struct{
	swarmid string
	host string
	updateid string
	//The name of the file
	Filename string
	//The filesize of the file
	Filesize uint64
	//The number of chunks which make up the file
	num_chunks uint16
	//The type of announcement which is being made, E. G. RequestFile
	AnnounceType int
	//Empty if requesting a file, otherwise it is pregnant with
	//the chunk to be retrieved or the content of said chunk.
	content []byte
}

func (f *FileAnnounce) SwarmId() string{
	return f.swarmid
}

func (f *FileAnnounce) UpdateId() string{
	return f.updateid
}
func (f *FileAnnounce) Type() string{
	return "FileAnnounce"
}
func (f *FileAnnounce) MarshalString() string{
	w,err:=json.Marshal(f)
	if err!=nil{
		panic("Unable to marshal
	FileAnnounceTransaction. (This shouldn't happen)"+err.Error())
	}
	return string(w)
}
