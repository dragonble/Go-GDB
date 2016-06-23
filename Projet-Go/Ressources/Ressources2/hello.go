package main

import (
 	"fmt"
	"github.com/cyrus-and/gdb"
	"io"
	"os"
)

type bidon struct {
	pos int32
}

func (bid *bidon) Write(p []byte) (n int, err error) {
	bid.pos++
	fmt.Println(bid.pos,p,string(p))
	return len(p),nil
}

func main() {
	bid := &bidon{}
	debug, _ := gdb.New(nil)
	debug.Send("file-exec-and-symbols",  os.Args[1])
	debug.Send("exec-run")
	io.Copy(bid, debug)
	debug.Exit()
}

