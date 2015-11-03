package mysocket

import (
	"fmt"
	"io"
	"math"
	"net"
	"strconv"
	"time"
)

func ServerGo() {
	defer WG.Done()
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		fmt.Printf("Listen Error %s\n", err)
		return
	}

	defer listener.Close()
	fmt.Printf("Got listenter for the server. (local address: %s)\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept Error %s\n", err)
		}
		fmt.Printf("Established a connection with a client application.(remote address: %s\n", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("The connection is closed by another side. (Server)")
			} else {
				fmt.Printf("Read Error: %s (Server)\n", err)
			}
			break
		}
		fmt.Printf("Received request: %s (Server)\n", strReq)

		i32Req, err := strconv.Atoi(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			if err != nil {
				fmt.Printf("Write Eoor (writen %d bytes:) %s (Server)\n", err)
			}
			fmt.Printf("Sent response (written %d bytes %s (server)\n", n, err)
			continue
		}

		f64Resp := math.Cbrt(float64(i32Req))
		respMsg := fmt.Sprintf("The cube root of %d is %f.", i32Req, f64Resp)
		n, err := write(conn, respMsg)
		if err != nil {
			fmt.Printf("Write Error: %s (Server)\n", err)
		}
		fmt.Printf("Sent response (writtn %d bytes:) %s (Server)\n", n, respMsg)
	}
}
