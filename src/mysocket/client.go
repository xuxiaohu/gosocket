package mysocket

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"
)

func ClientGo(id int) {

	defer WG.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		fmt.Printf("Dial Error: %s (Client[%d])\n", err, id)
		return
	}
	defer conn.Close()
	fmt.Printf("Connected to server. (remote address: %s, local address: %s) (Client[%d]\n", conn.RemoteAddr(), conn.LocalAddr(), id)
	time.Sleep(200 * time.Millisecond)

	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))

	for i := 0; i < requestNumber; i++ {
		i32Req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", i32Req))
		if err != nil {
			fmt.Printf("Write Error: %s (Client[%d]", err, id)
			continue
		}
		fmt.Printf("Send request (written %d bytes): %d (Client[%d])\n", n, i32Req, id)
	}

	for j := 0; j < requestNumber; j++ {
		strResq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("The connection is closed by another side. (Client[%d]\n", id)
			} else {
				fmt.Printf("Read Error: %s (Client[%d]\n", err, id)
			}
			break
		}
		fmt.Printf("Received response: %s (Client[%d]\n", strResq, id)
	}
}
