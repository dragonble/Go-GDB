package main

import (
	  "fmt"
	 "github.com/cyrus-and/gdb"
	  "io"
	  "os"
	 "strconv"
	  "path/filepath"
	"log"
)
	var debug *gdb.Gdb //Creation de la variable qui contiendra l'instance de GDB
	

	//Fonction qui demarre l'instance de GDB avec le fichier donne en arguments		
	func execGdb() {
		

				
						
					  debug, _ = gdb.New(nil) //Initialisation de l'instance de GDB
					 
					go io.Copy(os.Stdout, debug) // on affiche la sortie de GDB dans la console
						
					
					dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
   					 if err != nil {
            					log.Fatal(err)
    					}

					dirTravail := dir + "/Ressources"

			

					 
					  debug.Send("file-exec-and-symbols", dirTravail + "/" + os.Args[2]) // on donne le fichier compile a GDB
						
	}

	//Lance l'execution du programme avec creation du record pour pouvoir revenir en arriere si necessaire



	func start(){
				
				debug.Send("exec-run")				
				debug.Send("interpreter-exec","console","record")
				
				

	}
	//Avance d'une ligne
	func step(){
		
		 debug.Send("exec-step")
			
		
	}
	//Affiche la liste des breakpoints
	func breaklist(){

			output,err :=debug.Send("break-list") // On prends la reponse complexe de GDB
						if err !=nil {
							fmt.Println(err)				
							}

						pay:=output["payload"] //On va chercher etape par etape l'information voulue
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
							fmt.Println("number:",number,"type:",typeB,"enabled:",enabled,
							"times:",times,"disp:",disp,"function:",fun,"line:",line)
						}


		}

	//Supprime un breakpoint
	func delete_break(bpdel int){

				
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
							line :=bkptAssert["line"]
							ligneAssert:= line.(string)
							numLigne,_ := strconv.Atoi(ligneAssert)						
							if bpdel == numLigne{
								
								
								numero_break := number.(string)
								
								debug.Send("break-delete",numero_break )
								
								
							}
							
						}

						
																					  
							

	}
	
	//Recule d'une ligne
	func step_reverse(){
		debug.Send("exec-step","--reverse")

		}
	//Reprend l'execution 
	func continuee(){
			debug.Send("exec-continue")
		
		}
	//Revient au dernier breakpoint 
	func continue_reverse(){
			debug.Send("exec-continue","--reverse")

								
		}

	//Affiche la pile d'appels

	func backtrace() string {

			
				
								output,_:=debug.Send("stack-list-frames")
								pay:=output["payload"]
								
								payAssert:=pay.(map[string]interface{})
				
								stack:=payAssert["stack"]
							
								stackAssert:=stack.([]interface{})
								nbreFct:=len(stackAssert)
			
								str :=""
								str_variables := ""

								for i:=0; i<=nbreFct-1 ; i++{
									stackSepare:=stackAssert[i]
									stackSepareAssert:=stackSepare.(map[string]interface{})
								
									frame:=stackSepareAssert["frame"]
									frameAssert:=frame.(map[string]interface{})
									
									index := strconv.Itoa(i)

									
									//liste des variables par frame 
									output_variables,_  := debug.Send("stack-list-variables","--thread","1","--frame",index,"--simple-values")
									map_variables := output_variables["payload"]
									m_variables := map_variables.(map[string]interface{})

									variables := m_variables["variables"]
									//fmt.Println("Variables : ", variables)
									

									
									//variables := m_variables["variables"]
									
									
									variables_slice := variables.([]interface{})
									str_variables = ""
								
								for j:=0; j<len(variables_slice); j++ {
										Separe_variables := variables_slice[j]
										Separe_variablesAssert :=Separe_variables.(map[string]interface{})
										
										//Nom Variable 
										name := Separe_variablesAssert["name"]
										str_name := name.(string)
										
										//Type Variables 
										typee := Separe_variablesAssert["type"]
										str_type := typee.(string)
										
										//Valeur Variables 
										value := Separe_variablesAssert["value"]
										str_value := value.(string)
										

										// Verifie si la variable est un argument
										if arg, ok := Separe_variablesAssert["arg"] ;ok{
										str_arg := arg.(string)
											str_variables +="variable name : " + str_name + " type :" + str_type + " value :" + str_value +
																" arg : " + str_arg + "\n"
										}	else {
											str_variables +="variable name : " + str_name + " type :" + str_type +" value :" + str_value+ "\n"
											}
										
									}
								

									fun:=frameAssert["func"]
									line:=frameAssert["line"]
									level:=frameAssert["level"]
									

									function := fun.(string)
									ligne := line.(string)
									niveau := level.(string)
								
									str += "Frame " +index + "\n"+"level : " + niveau + " function : "+  function + "  line : " + ligne +"\n" + 														str_variables + "\n"
									

									
								}
			return str
		}


	//Creer un watchpoint

	func watch(){

				var input_watch string
				
				fmt.Println("Rentrez la variable suivi")
				fmt.Scanln(&input_watch)
				fmt.Println(debug.Send("break-watch", input_watch))
		}
	//Affiche ou en est l'execution du programme
	func where(){

			output,_ := debug.Send("stack-list-frames")
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
				if fun == "_exit" {
				fmt.Println("exit")
				debug.Exit()
				
				}
		}
	//Ajoute un breakpoint
	func breake(bp int){
			
						
			input_break := strconv.Itoa(bp)
			debug.Send("break-insert", input_break)
		}
	//Demande un nom de variable pour copier la valeur de la variable cible puis l'affiche
	func print(){

			var var_gdb string
			var var_cible string
		
			fmt.Println("Entrez le nom de la variable")
			fmt.Scanln(&var_gdb)
		
			fmt.Println("rentrez la variable cible")
			fmt.Scanln(&var_cible)
		
			debug.Send("var-create", var_gdb, "@", var_cible)
			output,err := debug.Send("var-evaluate-expression", var_gdb)	
				if err !=nil{
								fmt.Println(err)
						}

			fmt.Println(output["payload"])
	//Affiche la liste des variables	
		}
	func list_variables (){
			
			expr,err := debug.Send("stack-list-variables", "--all-values")	
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
	//Arrete GDB

	func stop(){
		 debug.Interrupt()
		 debug.Exit()
	}
	


