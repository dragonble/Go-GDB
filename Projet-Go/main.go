	package main

	import (
	    "fmt"
	    "github.com/cyrus-and/gdb"
	    "io"
	    "os"
	)

	func main() {
	    // start a new instance and pipe the target output to stdout
	    gdb, _ := gdb.New(nil)
	    go io.Copy(os.Stdout, gdb)

	    // evaluate an expression
	    gdb.Send("var-create", "x", "@", "40 + 2")
	    fmt.Println(gdb.Send("var-evaluate-expression", "x"))

	    // load and run a program
	    gdb.Send("file-exec-and-symbols", "prog")
	    gdb.Send("exec-arguments", "-w")
	    gdb.Write([]byte("This sentence has five words.\n\x04")) // EOT
	    //gdb.Send("exec-run")
	    //gdb.Interrupt()
	    gdb.Write([]byte("Entrez votre breakpoint.\n\x04"))   
	    var input string
	    fmt.Scanln(&input)
	    gdb.Send("break-insert",input)
	    gdb.Send("exec-run")
	    /*err:= gdb.Interrupt()
	if err != nil {
	    fmt.Println(err)
	    }*/
	    gdb.Send("var-create", "c_name", "@", "c")
	   fmt.Println( gdb.Send ("var-evaluate-expression", "c_name"))
	    gdb.Exit()
	}
