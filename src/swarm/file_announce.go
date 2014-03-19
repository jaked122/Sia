package swarm

import (
	"encoding/json"
)

const (
	//Request a file from the network
	RequestFile = iota
	//Serve a file for the network to obtain
	ServeFile
	//Request a particular chunk to be stored
	RequestChunk
	//Delete the file
	DeleteFile
)

//Stores the necessary information to make the file transmittable
type FileAnnounce struct {
	Origin     string
	Filename   string
	FileSize   uint64
	Chunk      uint16
	swarmid    string
	updateid   string
	Content    []byte
	recordtype int
}

func NewFileAnnounce(Requesttype int, chunk uint16, origin, filename, swarmid string) *FileAnnounce {
	i := new(FileAnnounce)
	i.Origin = origin
	i.Chunk = chunk
	i.recordtype = Requesttype
	i.swarmid = swarmid
	return i
}
func (f *FileAnnounce) SwarmId() string {
	return f.swarmid
}

func (f *FileAnnounce) UpdateId() string {
	return f.updateid
}

func (f *FileAnnounce) Type() string {
	return "FileAnnounce"
}
func (f *FileAnnounce) MarshalString() string {
	w, err := json.Marshal(f)
	if err != nil {
		panic("Unable to marshal FileAnnounceTransaction. (This shouldn't happen)" + err.Error())
	}
	return string(w)
}
