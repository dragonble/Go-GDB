package main

import (
 	"fmt"
	"github.com/cyrus-and/gdb"
	"io"
	"os"
)

type outputGDB struct {
	pos int32
}

func (out *outputGDB) Write(p []byte) (n int, err error) {
	out.pos++
	fmt.Println(string(p))
	return len(p),nil
}

func main() {
	out := &outputGDB{}
	debug, _ := gdb.New(nil)
	debug.Send("file-exec-and-symbols",  os.Args[1])
	debug.Send("exec-run")
	io.Copy(out, debug)
	fmt.Println("pas normal")
	//debug.Exit()
	fmt.Println("pas normal")
	//os.Exit(0)
	
}

