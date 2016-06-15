package main

			import (
					  "fmt"
					  "github.com/cyrus-and/gdb"
					  "io"
					  "os"
					  //"path/filepath"
					"strconv"
				)
					var debug *gdb.Gdb
				
					func main() {
			  // start a new instance and pipe the target output to stdout
					  debug, _ = gdb.New(nil)
					  go io.Copy(os.Stdout, debug)

			

					  // load and run a program
					  debug.Send("file-exec-and-symbols", os.Args[1])
					fmt.Println(debug.Send("gdb-version"))

					 
					  var input string
					  
				 for input != "quit"   {
					  fmt.Println("Que voulez-vous faire ?")
					  fmt.Scanln(&input)
					
						switch input {
						
						//Break
						case "break": breake()
						
						//Run	
						case "run" : start()
						

						//Backtrace

						case "backtrace" :backtrace()
	

									// Default Case					
						default: fmt.Println("Commandes non valides")  

					}
					  
				}	

					  debug.Exit()
	}


	func start() (){
					
				fmt.Println(debug.Send("exec-run"))
				debug.Send("interpreter-exec","console","record")	

			}

	func breake(){

			var input_break string
			fmt.Println("rentrez votre breakpoint")
			fmt.Scanln(&input_break)
			debug.Send("break-insert", input_break)
		}

	func backtrace(){
			
				
								output,_:=debug.Send("stack-list-frames")
								pay:=output["payload"]

								payAssert:=pay.(map[string]interface{})
				
								stack:=payAssert["stack"]
							
								stackAssert:=stack.([]interface{})
								nbreFct:=len(stackAssert)
								for i:=0; i<=nbreFct-1 ; i++{
									stackSepare:=stackAssert[i]
									stackSepareAssert:=stackSepare.(map[string]interface{})
								
									frame:=stackSepareAssert["frame"]
									frameAssert:=frame.(map[string]interface{})
									
									index := strconv.Itoa(i)
									fmt.Println(index)
									//list variables by frame 
									fmt.Println(debug.Send("stack-list-variables","--thread","1", "--frame",index,"--simple-values"))

									fun:=frameAssert["func"]
									line:=frameAssert["line"]
									level:=frameAssert["level"]
									fmt.Println("level : ",level,"function : ",fun ,"  line : ",line)
									
									
								}
		}


		


