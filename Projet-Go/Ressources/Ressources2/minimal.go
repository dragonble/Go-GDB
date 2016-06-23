package main

import (
 	"fmt"
	"github.com/cyrus-and/gdb"
	"io"
	"os"
	 "io/ioutil" 
)

func main() {
	
	
	debug, _ := gdb.New(nil) //Initialisation de l'instance de GDB
	go io.Copy(os.Stdout, debug) // on affiche la sortie de GDB dans la console
	debug.Send("file-exec-and-symbols",  os.Args[1])

	rescueStdout := os.Stdout
 	 r, w, _ := os.Pipe()
 	 os.Stdout = w

 
	debug.Send("exec-run")
	
	
	fmt.Println("Hello, playground") 
  	w.Close()
  	out, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout

	 fmt.Printf("Captured: %s \n", out)
	debug.Exit()

}












