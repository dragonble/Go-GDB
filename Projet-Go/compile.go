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
	    gdb.Send("file-exec-and-symbols", "essai")
	   
	   /* gdb.Send("break-insert" , "6" )


	   
	    

	    gdb.Send("exec-run" )
	    gdb.Send("var-create", "cname", "@", "c")
	    express, err := gdb.Send("var-evaluate-expression", "cname")	
	    if err != nil{
			fmt.Println(err)
		}
	    fmt.Println(express)*/
	    var input string
	    
	    for input != "quit"   {
	    fmt.Println("Que voulez-vous faire ?")
	    fmt.Scanln(&input)
	    
		//Continue
	   if input == "continue" {
	    gdb.Send("exec-continue")
	    }

		//Step
	    if input == "step" {
	   fmt.Println( gdb.Send("exec-step"))
	    }

		//
	   /* if input == "print" {
	    gdb.Send("var-create", "cname", "@", "c")
	    gdb.Send("var-evaluate-expression", "cname")
	    }*/
		
		//BREAK
	    if input == "break" {
		var input_break string
		fmt.Println("rentrez votre breakpoint")
		fmt.Scanln(&input_break)
		gdb.Send("break-insert", input_break)
		}
		
		//Breakpoints list
	    if input == "break-list" {
		fmt.Println(gdb.Send("break-list"))
		}

		//Print
	    if input == "print" {
		
		var var_gdb string
		var var_cible string
		
		fmt.Println("Entrez le nom de la variable")
		fmt.Scanln(&var_gdb)
		
                fmt.Println("rentrez la variable cible")
		fmt.Scanln(&var_cible)
		
		gdb.Send("var-create", var_gdb, "@", var_cible)
		express2,err := gdb.Send("var-evaluate-expression", var_gdb)	
		if err !=nil{
			fmt.Println(err)
			}
		fmt.Println(express2)
		}
		
		//Run
	    if input == "run" {
		gdb.Send("exec-run" )		
		}


	    

	        //Variables List locals
		
	    if input == "list" {		
		fmt.Println(gdb.Send("stack-list-variables","1"))
		}


	}
	    //gdb.Exit()
}


