package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/unicornzyn/gotest27/common/message"
)

// Transfer 数据传输结构体
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

// ReadPkg 读消息
func (transfer *Transfer) ReadPkg() (msg message.Message, err error) {
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[0:4])

	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen]) err=", err)
		return
	}
	// 这里注意 &msg 传地址 否则msg作为结构体是值类型 结果不能带出来
	err = json.Unmarshal(transfer.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &msg) err=", err)
		return
	}
	return
}

// WritePkg 写消息
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	_, err = transfer.Conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.Write(pkgLen) err=", err)
		return
	}
	//fmt.Println("客户端发送消息的长度ok")

	_, err = transfer.Conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	return
}
