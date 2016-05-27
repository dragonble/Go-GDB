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
					 
					  var input string
					  
				 for input != "quit"   {
					  fmt.Println("Que voulez-vous faire ?")
					  fmt.Scanln(&input)
					  
					//Continue
					 if input == "continue" {
					  fmt.Println(gdb.Send("exec-continue"))
					  }

					//Reverse continue 
					if input =="continue-reverse"{
						fmt.Println(gdb.Send("exec-continue","--reverse"))
				}

					//Step
					  if input == "step" {
					  output, err := gdb.Send("exec-step")
					if err != nil {
					fmt.Println(err)		
					}
					fmt.Println(output)
			}
					//Reverse stepping
					if input == "step-reverse" {
							fmt.Println(gdb.Send("exec-step","--reverse"))
					
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
					gdb.Send("interpreter-exec","console","record")	
					}

				
						  //List of Variables locals
					  if input == "list-variables" {

				
						expr,err := gdb.Send("stack-list-variables", "--all-values")	
						if err !=nil {
							fmt.Println(err)
						}

						variables := expr["payload"]
				
						variablesAssert := variables.(map[string]interface {})
						Array := variablesAssert["variables"]
						ArrayAssert := Array.([]interface{})
						nbreVar:=len(ArrayAssert)
						for i:=0; i<=nbreVar-1 ; i++{
							mapListe := ArrayAssert[i]
							mapListeAssert := mapListe.(map[string]interface {})
							name := mapListeAssert["name"]
							value := mapListeAssert["value"]
							arg := mapListeAssert["arg"]
							fmt.Println("name : ", name , "value : ",value , "arg : ",arg)
								
						}
				

					}


				
							//Backtrace
						  if input == "backtrace"{

								output,_:=gdb.Send("stack-list-frames")
								pay:=output["payload"]

								/*expr,err:=gdb.Send("stack-list-frames")
								if err != nil {
									fmt.Println(err)		
								}
								pay:= expr["payload"]*/

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
							
							

						} 

					//Watcpoints
					if input == "watch" {
						var input_watch string
						fmt.Println("Rentrez la vairable suivi")
						fmt.Scanln(&input_watch)
						fmt.Println(gdb.Send("break-watch", input_watch))

					}
						
					
					
					//Where
				if input == "where" {
						output,_ := gdb.Send("stack-list-frames")
							pay:=output["payload"]

							payAssert:=pay.(map[string]interface{})
				
							stack:=payAssert["stack"]
							stackAssert:=stack.([]interface{})
							
							//Premier stack			
							stackSepare:=stackAssert[0]
							stackSepareAssert:=stackSepare.(map[string]interface{})
								
							frame:=stackSepareAssert["frame"]
							frameAssert:=frame.(map[string]interface{})
							fun:=frameAssert["func"]
							line:=frameAssert["line"]
							fmt.Println("function : ",fun ,"  line : ",line)

				}
 	}
	

					  gdb.Exit()
				}



