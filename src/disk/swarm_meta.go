package disk

import "strconv"

type FilePreferences struct{
	//The number of chunks in the file
	ChunkCount int
	//The number of times a chunk should be duplicated
	DesiredChunkDuplication int
}
type SwarmMeta struct{
	//The storage associated
	Storage *SwarmStorage
	//Stores the number of chunks for each file on the "Blockchain"(Swarm)
	FileChunking map[string]FilePreferences
	//Stores which chunk is available for the "Blockchain"(Swarm) that are 
	//stored in this node
	AvailableChunks map[string]int
}

func NewSwarmMeta(sname string) *SwarmMeta{
	r:=new( SwarmMeta)
	r.Storage,_=CreateSwarmSystem(sname)
	r.FileChunking=make(map[string]FilePreferences)
	r.AvailableChunks=make(map[string]int)
	return r
}
//Create A file
func (r SwarmMeta) CreateFile(N string,chunkid int, i []byte){
	fname:=string(strconv.AppendInt([]byte(N),int64(chunkid),10))
	r.Storage.CreateFile(fname,uint64(len(i)))
	r.Storage.WriteFile(fname,0,i)
}
func (r SwarmMeta) ReadFile(N string, chunkid int) []byte{
    f:=int64(chunkid)
	c:=make([]byte,r.Storage.files[string(strconv.AppendInt([]byte(N),f,10))])
	err:=r.Storage.ReadFile(string(strconv.AppendInt([]byte(N),f,10)),0, c)
	if(err!=nil){
	}
	return c
}
func (r SwarmMeta) IsThereSpaceLeft()bool{
	return r.Storage.amountused<(2<<32)
}
