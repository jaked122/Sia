package swarm

import "encoding/json"
import "strconv"

type FileResponse struct{
    Filename string
    FileSize uint64
    Swarmid string
    Response_Origin string
    Chunk uint16
    Response_ID uint64
}
func NewFileResponse(content []byte,Filename, Response_Origin, swarmid,update string,
filesize,response_id uint64,c uint16) (i *FileResponse){
    i=new(FileResponse)
    i.Filename=Filename
    i.FileSize=filesize
    i.Response_Origin=Response_Origin
    i.Swarmid=swarmid
    i.Response_ID=response_id
    i.chunk=c
    i.updateid=update
}
func (f *FileResponse) Type()string{
    return "FileResponse"
}
//People in general love hexadecimal
func (f *FileResponse) UpdateId()string{
    return strconv.FormatUint(f.Response_ID,16)
}
func (f *FileResponse) MarshalString()string{
    w,err:=json.Marshal(f)
    if err!=nil{
        panic("Unable to marshal FileResponse")
    }
    return string(w)
}
