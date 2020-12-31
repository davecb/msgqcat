package main

import (
	"fmt"
	"github.com/siadat/ipc"
	"log"
	"os"
	"strconv"
)

func main() {
	var qID uint       // queue ID
	var mtype uint = 0 // any type
	var mtext []byte = make([]byte, 10000)
	var buff = ipc.Msgbuf{mtype, mtext}
	var flags uint = ipc.IPC_NOWAIT // return ENOMSG on EOF

	if len(os.Args) != 2 {
		log.Fatalf("Usage: msgqcat queue_id\n")
	}
	q := os.Args[1]
	id, err := strconv.ParseUint(q, 10, 64)
	qID = uint(id)
	if err != nil {
		log.Fatalf("msgqcat: the queue_id to read from must be an integer")
	}

	for i := 0; ; i++ {
		err = ipc.Msgrcv(qID, &buff, flags)
		if err != nil {
			if err.Error() == "no message of desired type" {
				// That's an EOF is ipc-ish
				break
			}
			panic(fmt.Errorf(`ipc.Msgrcv failed with the message "%v"`, err))
		}
		fmt.Printf("%s\n", buff)
	}
}
