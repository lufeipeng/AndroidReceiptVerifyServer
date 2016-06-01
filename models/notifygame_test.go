package models

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"net"
	"testing"
)

type BillCallbackMsg struct {
	Uid   int `json:"uid"`
	Serid int `json:"serid"`
}

type BillCallbackInfo struct {
	MsgCmd int             `json:"msg_cmd"`
	Sid    int             `json:"sid"`
	Msg    BillCallbackMsg `json:"msg"`
}

func PackMsg(uid, sid, serid int) []byte {
	var billinfo BillCallbackInfo
	billinfo.MsgCmd = 1201
	billinfo.Sid = (4<<16 | sid)
	billinfo.Msg.Serid = serid
	billinfo.Msg.Uid = uid

	outByte, err := json.Marshal(billinfo)
	if err != nil {
		fmt.Println("PackMsg is call error")
		return []byte{}
	}
	return outByte
}

func TestNet(t *testing.T) {
	conn, err := net.Dial("tcp", "61.132.227.32:52201")
	if err != nil {
		fmt.Println("TestXYZ is called")
		return
	}
	msgBuffer := PackMsg(34022, 1, 1035)
	fmt.Println("TestNet msg is ", string(msgBuffer))

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint32(len(msgBuffer)))
	binary.Write(buf, binary.BigEndian, uint16(1002))
	binary.Write(buf, binary.BigEndian, uint16(0))
	binary.Write(buf, binary.BigEndian, msgBuffer)
	writeLen, err := conn.Write(buf.Bytes())
	fmt.Println("TestNet is called", writeLen, " : ", err)

	var num int32
	readBuffer := make([]byte, 4)
	length, err := conn.Read(readBuffer)
	if err != nil {
		fmt.Println("TestNet is read", err)
		return
	}
	fmt.Println("TestNet is read, pack size is", length)
	newReadBuffer := bytes.NewReader(readBuffer)
	err = binary.Read(newReadBuffer, binary.BigEndian, &num)
	fmt.Println("TestNet is read, size is ", num)
}
