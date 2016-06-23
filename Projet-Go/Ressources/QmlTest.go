package main

import (
    "fmt"
    "gopkg.in/qml.v1"
    "os"
    "io/ioutil"    
    "path/filepath"
    "log"
  //  "time"
    "io"
   // "bytes"
	//"os/exec"
	
)

func main() {

	


    if err := qml.Run(run); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
   	 if err != nil {
            log.Fatal(err)
    	}

	dirTravail := dir + "/Ressources2"

    dat, err2 := ioutil.ReadFile(dirTravail + "/" + os.Args[1])
    if err2 != nil {
            log.Fatal(err2)
    	}

    execGdb()
   
    engine := qml.NewEngine()

    component, err := engine.LoadFile(dir + "/fenetre.qml")
    if err != nil {
        return err
    }

    context := engine.Context()
    context.SetVar("fileOp",&FileTest{Content : string(dat)})
<<<<<<< HEAD

	
    //context2 := engine.Context()
    context.SetVar("consoleOp",&FileTest{Content : ""})
	


=======
	
>>>>>>> e71d172bf9a7910559129adadfa51b122c4af6c9
    window := component.CreateWindow(nil)
	

    window.Show()
    window.Wait()

    return nil
}

type FileTest struct {
    Name    string
    Content string
}

<<<<<<< HEAD
func (consoleOp *FileTest) Debugrun() {  
	/*rescueStdout := os.Stdout
 	 r, w, _ := os.Pipe()
 	 os.Stdout = w*/

 
	start()
	io.Copy(os.Stdout, debug)
	
	fmt.Println("Hello, playground") 
	debug.Exit()
  	/*w.Close()
	debug.Interrupt()
  	out, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout*/

 
	



	 //fmt.Printf("Captured: %s \n", out)
	//debug.Interrupt()
  	//consoleOp.Content = consoleOp.Content + "\n" + string(out)
	//qml.Changed(consoleOp, &consoleOp.Content)
	//debug.Exit()
	
=======
func (fileOp *FileTest) Debugrun() {  
        start() 
	fmt.Println("start")
	where()
	
	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
>>>>>>> e71d172bf9a7910559129adadfa51b122c4af6c9
}
func (fileOp *FileTest) Debugstep() {  
	
        step()
	fmt.Println("step") 
	where()
	
<<<<<<< HEAD
=======
	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
>>>>>>> e71d172bf9a7910559129adadfa51b122c4af6c9
}
func (fileOp *FileTest) Debugcontinue() {  
	
        continuee() 
	fmt.Println("continue")
<<<<<<< HEAD
	where()
=======

	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
>>>>>>> e71d172bf9a7910559129adadfa51b122c4af6c9
	
}

func (fileOp *FileTest) Debugreverse() {  
        debug.Exit()
	//step_reverse()
	fmt.Println("reverse")
	
	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
}

func (fileOp *FileTest) Addbreakpoint(bp int) {  
        
	
	breake(bp)
	fmt.Println(bp)
}
func (fileOp *FileTest) Rmvbreakpoint(bp int){
	delete_break(bp)
	fmt.Println(bp)
}
/*
func (consoleOp *FileTest) Affichebuffer() {

<<<<<<< HEAD
	
	 rescueStdout := os.Stdout
 	 r, w, _ := os.Pipe()
  	os.Stdout = w
=======
func (fileOp *FileTest) Debugstop(){
	stop()
}	


<<<<<<< HEAD
func (fileOp *FileTest) Debugreversecontuinue(){
	continue_reverse()
	
	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
}
=======
>>>>>>> e71d172bf9a7910559129adadfa51b122c4af6c9

>>>>>>> 17e46f2ad7e31fd4a9ca13e71ed2772c402360f9

  	w.Close()
  	out, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout
	
  	consoleOp.Content = consoleOp.Content + "\n" + string(out)
	qml.Changed(consoleOp, &consoleOp.Content)
}


*/



