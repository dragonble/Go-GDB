package main

import (
    "fmt"
    "gopkg.in/qml.v1"
    "os"
    "io/ioutil"    
    "path/filepath"
    "log"
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
	
    window := component.CreateWindow(nil)
	

    window.Show()
    window.Wait()

    return nil
}

type FileTest struct {
    Name    string
    Content string
}

func (fileOp *FileTest) Debugrun() {  
        start() 
	fmt.Println("start")
	where()
	
	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
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

	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
	
}

func (fileOp *FileTest) Debugreverse() {  
        
	step_reverse()
	fmt.Println("reverse")
	//Backtrace
	fileOp.Name = backtrace()
	qml.Changed(fileOp, &fileOp.Name)
}

func (fileOp *FileTest) Addbreakpoint(bp int) {  
        
	
	breake(bp)
}
func (fileOp *FileTest) Rmvbreakpoint(bp int){
	fmt.Println(bp)
}

func (fileOp *FileTest) Debugstop(){
	stop()
}	









