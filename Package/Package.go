package Package

import (
	"ReAce/Utils"
	"bytes"
)

type Package struct {
	uuid   string
	cmd    int
	params any
}

/*
*	COMMAND ID
*
 */
const (
	HEART = iota
	LOGIN_RECV
	SHELL_SENT
	SHELL_RECV

	SHELLCODE_LOAD

	PROC_GETLIST
	PROC_RECVLIST
	FS_GETFILES
	FS_RECVFILES
	FS_CD
	FS_CURRENT
	FS_DOWNLOAD
	FS_ERROR
	FS_SEND
)

func NewPackage(uuid string, cmd int, params any) *Package {
	return &Package{
		uuid:   uuid,
		cmd:    cmd,
		params: params,
	}
}

// getter
func (p *Package) GetUUID() string {
	return p.uuid
}
func (p *Package) GetCmd() int {
	return p.cmd
}
func (p *Package) GetParams() any {
	return p.params
}
func GetPackLen(pack []byte) int {
	if bytes.Compare(pack[:2], []byte{0xde, 0xad}) != 0 || len(pack) < 6 {
		return 0
	}

	return Utils.BytesToInt(pack[2:6])
}

func GetPackFromBytes(pack []byte) *Package {
	//fmt.Println(pack)
	tempUuid := string(pack[:10])
	tempCmd := Utils.BytesToInt(pack[10:14])
	tempParam := pack[14:]
	x := NewPackage(tempUuid, tempCmd, tempParam)
	//fmt.Println("uuid:" + tempUuid)
	//fmt.Println("cmd:", tempCmd)
	//fmt.Println("其他:", tempParam)
	return x
}

// 数据包构造
func (p *Package) MakePackage() []byte {

	res := make([]byte, 6)
	res = append(res, []byte(p.uuid)...)
	res = append(res, Utils.IntToBytes(p.cmd)...)
	if params, ok := p.params.([]string); ok {
		//tmpLen := len(params)
		for _, v := range params {
			res = append(res, []byte(v)...)
			res = append(res, '\x00')
		}

	} else if param, ok := p.params.(string); ok {
		res = append(res, []byte(param)...)
		res = append(res, '\x00')
	} else if param, ok := p.params.([]byte); ok {
		res = append(res, []byte(param)...)
	}
	res[0] = 0xde
	res[1] = 0xad
	pLen := Utils.IntToBytes(len(res) - 6)
	for i := 0; i < 4; i++ {
		res[i+2] = pLen[i]
	}
	//fmt.Println(res)
	return res
}
