package main

import (
	. "mysocket"
	"time"
)

// func main() {
// 	go serverGo()
// 	time.Sleep(500 * time.Millisecond)
// 	go clientGo(1)
// }

func main() {
	WG.Add(2)
	go ServerGo()
	time.Sleep(500 * time.Millisecond)
	go ClientGo(1)
	WG.Wait()
}
