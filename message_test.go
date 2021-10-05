package simpleTcpMessage

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"sync"
	"testing"
)

const TEST_IP = "127.0.0.1:25565"
const TEST_PORT = "25565"

func TestSendLongMessage(t *testing.T) {
	wg := sync.WaitGroup{}
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Add(1)
	sendArr, _ := os.ReadFile("test_files/large_text.txt")
	go func() {
		conn, err := net.Dial("tcp", "127.0.0.1:9999")
		if err != nil {
			fmt.Println(err.Error())
		}
		clientConn := NewClientConn(conn)
		msg := NewMessage()

		msg.AppendField(0, sendArr)
		msg.AppendField(1, sendArr)
		msg.AppendField(2, sendArr)
		// msg.AppendField(1, []byte("Test Message"))
		clientConn.SendMessage(msg)

		wg.Done()
	}()

	conn, _ := ln.Accept()
	clientConn := NewClientConn(conn)

	msg, _ := clientConn.RecieveMessage()
	recieveArr1, _ := msg.GetField(0)
	recieveArr2, _ := msg.GetField(1)
	recieveArr3, _ := msg.GetField(2)

	ok1 := reflect.DeepEqual(sendArr, recieveArr1)
	ok2 := reflect.DeepEqual(sendArr, recieveArr2)
	ok3 := reflect.DeepEqual(sendArr, recieveArr3)

	if !ok1 || !ok2 || !ok3 {
		t.Fail()
	}
	wg.Wait()
}
