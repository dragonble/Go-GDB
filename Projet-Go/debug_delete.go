package main

				import (
					  "fmt"
					  "github.com/cyrus-and/gdb"
					  "io"
					  "os"
					  //"path/filepath"
				)
					var debug *gdb.Gdb
				func main() {


					  // start a new instance and pipe the target output to stdout
					  debug, _ = gdb.New(nil)
					  go io.Copy(os.Stdout, debug)

			

					  // load and run a program
					  debug.Send("file-exec-and-symbols", os.Args[1])
					 
					  var input string
					  
				 for input != "quit"   {
					  fmt.Println("Que voulez-vous faire ?")
					  fmt.Scanln(&input)
					if input == "delete" {
								delete_break()							
							}
					if input == "break"{
								breake()	
							}
						if input ==  "break-list" {
								breaklist()
						}
					}
					

	}


	func delete_break(){
				
						var numero_break string		
				
						fmt.Println("Supprimer un breakpoint(n°) ou tous les breakpoints")
						fmt.Scanln(&numero_break )
						
						
						fmt.Println(debug.Send("break-delete",numero_break))
						
							

	}

	func breake(){

			var input_break string
			fmt.Println("rentrez votre breakpoint")
			fmt.Scanln(&input_break)
			debug.Send("break-insert", input_break)
		}

	func breaklist(){

			output,err :=debug.Send("break-list")
						if err !=nil {
							fmt.Println(err)				
							}

						pay:=output["payload"]
						payAssert := pay.(map[string]interface {})
					
						breakpointTable := payAssert["BreakpointTable"]
						breakpointTableAssert := breakpointTable.(map[string]interface {})
					
						Array := breakpointTableAssert["body"]
						ArrayAssert := Array.([]interface{})
						nbreVar:=len(ArrayAssert)
					
						for i:=0; i<=nbreVar-1 ; i++{
							mapSepare := ArrayAssert[i]
							mapSepareAssert := mapSepare.(map[string]interface {})
							bkpt := mapSepareAssert["bkpt"]
							bkptAssert := bkpt.(map[string]interface {})
							number:=bkptAssert["number"]
							typeB :=bkptAssert["type"]
							enabled :=bkptAssert["enabled"]
							times :=bkptAssert["times"]
							disp :=bkptAssert["disp"]
							fun :=bkptAssert["func"]
							line :=bkptAssert["line"]
							fmt.Println("number:",number,"type:",typeB,"enabled:",enabled,"times:",times,"disp:",disp,"function:",fun,"line:",line)
						}


		}
