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
)

//Stores the necessary information to make the file transmittable
type FileAnnounce struct {
	Origin      string
	Filename    string
	FileSize    uint64
	TotalChunks uint16
	Chunk       uint16
	swarmid     string
	updateid    string
	Content     []byte

}

//Create a new fileannounce which presumably will be sent across
// the wire to other hosts.
func NewFileAnnounce(filename, origin, swarmid, updateid string,
	TotalChunks, Chunk uint16, content []byte, blockchain *Blockchain) *FileAnnounce {
	i := new(FileAnnounce)
	i.Origin = origin
	i.Filename = filename
	i.FileSize = (uint64)(len(content)) * (uint64)(TotalChunks)
	i.TotalChunks = TotalChunks
	i.Content = content
	i.swarmid = swarmid
	i.updateid = updateid
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
