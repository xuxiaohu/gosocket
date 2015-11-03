package mysocket

import (
	"sync"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = "\t"
)

var WG sync.WaitGroup
