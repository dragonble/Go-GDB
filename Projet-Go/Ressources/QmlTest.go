package main

import (
    "fmt"
    "gopkg.in/qml.v1"
    "os"
    "io/ioutil"    
    "path/filepath"
    "log"
  //  "time"
    //"io"
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

    component, err := engine.LoadFile(dir + "/main.qml")
    if err != nil {
        return err
    }

    context := engine.Context()
    context.SetVar("fileOp",&FileTest{Content : string(dat)})

	
    //context2 := engine.Context()
    context.SetVar("consoleOp",&FileTest{Content : ""})

	//backtrace
	context.SetVar("backtraceFile", &FileTest{Content : ""})
	
    window := component.CreateWindow(nil)
	

    window.Show()
    window.Wait()

    return nil
}

type FileTest struct {
    Name    string
    Content string
}

/*<<<<<<< HEAD
func (consoleOp *FileTest) Debugrun() {  
	/*rescueStdout := os.Stdout
 	 r, w, _ := os.Pipe()
 	 os.Stdout = w*/

 
	/*start()
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
/*	
=======*/
func (backtraceFile *FileTest) Debugrun() {  
        start() 
	fmt.Println("start")
	where()
	
	//Backtrace
	backtraceFile.Content =backtraceFile.Content +	 backtrace()
	qml.Changed(backtraceFile, &backtraceFile.Content)

}
func (fileOp *FileTest) Debugstep() {  
	
        step()
	fmt.Println("step") 
	where()
	

	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)

}
func (fileOp *FileTest) Debugcontinue() {  
	
        continuee() 
	fmt.Println("continue")

	where()


	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)

	
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


func (fileOp *FileTest) Debugreversecontuinue(){
	continue_reverse()
	
	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
}

func (fileOp *FileTest) Debugstop(){
	stop()
}

/*
func (consoleOp *FileTest) Affichebuffer() {
	
	 rescueStdout := os.Stdout
 	 r, w, _ := os.Pipe()
  	os.Stdout = w

  	w.Close()
  	out, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout
	
  	consoleOp.Content = consoleOp.Content + "\n" + string(out)
	qml.Changed(consoleOp, &consoleOp.Content)
}


*/



