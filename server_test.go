package simpleTcpMessage

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

const TEST_IP = "127.0.0.1:25565"
const TEST_PORT = "25565"

func TestSend(t *testing.T) {
	wg := sync.WaitGroup{}
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Add(1)
	go func() {
		conn, err := net.Dial("tcp", "127.0.0.1:9999")
		if err != nil {
			fmt.Println(err.Error())
		}
		clientConn := NewClientConn(conn)
		msg := NewMessage()

		msg.AppendField(0, []byte("User1"))
		msg.AppendField(1, []byte("Test Message"))
		clientConn.SendMessage(msg)

		wg.Done()
	}()

	conn, _ := ln.Accept()
	clientConn := NewClientConn(conn)

	msg, _ := clientConn.RecieveMessage()

	t.Log(msg)
	wg.Wait()
}
