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
				  output, err := gdb.Send("exec-step")
				if err != nil {
				fmt.Println(err)		
				}
				fmt.Println(output)
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
					output,err :=gdb.Send("break-list")
					if err !=nil {
						fmt.Println(err)				
						}
					fmt.Println(output["payload"])
				}

				//Break delete
				if input == "delete" {
						var numero_break string		
						fmt.Println("Supprimer un breakpoint(n°) ou tous les breakpoints")
						fmt.Scanln(&numero_break )
						if numero_break != "" {
						gdb.Send("break-delete",numero_break )
					} else {
						gdb.Send("break-delete")
						}
		
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
				output,err := gdb.Send("var-evaluate-expression", var_gdb)	
				if err !=nil{
					fmt.Println(err)
					}

				fmt.Println(output["payload"])
				}
		
				//Run
				  if input == "run" {
				gdb.Send("exec-run" )		
				}

				
					  //List of Variables locals
				  if input == "list-variables" {

				fmt.Println(gdb.Send("stack-list-variables", "--all-values"))

				}

<<<<<<< HEAD
					//Backtrace
				if input == "backtrace"{
							fmt.Println(gdb.Send("stack-list-frames"))
=======
	    

			    /*//Variables List locals
		
						if input == "list" {		
							fmt.Println(gdb.Send("stack-list-variables","1"))
						}*/
	

						//Backtrace
					  if input == "backtrace"{
							expr,_:=gdb.Send("stack-list-frames")
							pay:=expr["payload"]
							payAssert:=pay.(map[string]interface{})
				
							stack:=payAssert["stack"]
							
							stackAssert:=stack.([]interface{})
							nbreFct:=len(stackAssert)
							for i:=0; i<=nbreFct-1 ; i++{
								stackSepare:=stackAssert[i]
								stackSepareAssert:=stackSepare.(map[string]interface{})
								
								frame:=stackSepareAssert["frame"]
								frameAssert:=frame.(map[string]interface{})
								fun:=frameAssert["func"]
								line:=frameAssert["line"]
								level:=frameAssert["level"]
								fmt.Println("level : ",level,"function : ",fun ,"  line : ",line)

								
							}
							
							
>>>>>>> dd734bb6975e77613a9c594d822ebb9aff62310d
					} 

				//Watcpoints
				if input == "watchpoints" {


				}

					//Record
				if input == "record" {
					fmt.Println(gdb.Send("record"))
			}
					//Reverse stepping
				if input == "step-reverse" {
					fmt.Println(gdb.Send("exec-step","--reverse"))
					


					}





}
	

				  gdb.Exit()
			}



