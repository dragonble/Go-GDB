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

  

    // load and run a program
    gdb.Send("file-exec-and-symbols", "essai.exe")
   
    gdb.Send("break-insert" , "6" )


   
    

    gdb.Send("exec-run" )
    gdb.Send("var-create", "cname", "@", "c")
    fmt.Println(gdb.Send("var-evaluate-expression", "cname"))
    var input string
    for input != "quit"   {
    fmt.Println("Que voulez-vous faire ?")
    fmt.Scanln(&input)
    if input == "continue" {
    gdb.Send("exec-continue")
    }
    if input == "step" {
    gdb.Send("exec-step")
    }
    if input == "print" {
    gdb.Send("var-create", "cname", "@", "c")
    fmt.Println(gdb.Send("var-evaluate-expression", "cname"))
    }
    }

   



    //gdb.Exit()
}


