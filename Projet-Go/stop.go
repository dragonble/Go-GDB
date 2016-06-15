package main

				import (
					  "fmt"
					  "github.com/cyrus-and/gdb"
					  "io"
					  "os"
)

		var debug *gdb.Gdb
				func main() {


					  // start a new instance and pipe the target output to stdout
					  debug, _ = gdb.New(nil)
					  go io.Copy(os.Stdout, debug)

			

					  // load and run a program
					  debug.Send("file-exec-and-symbols", os.Args[1])
					fmt.Println(debug.Send("gdb-version"))
					start()
					stop()

					 
}
	func start() {
					
				fmt.Println(debug.Send("exec-run"))
				debug.Send("interpreter-exec","console","record")	

			}

	func stop(){
		 debug.Interrupt()
		 debug.Exit()
}
