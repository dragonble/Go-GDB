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

	dirTravail := dir + "/Ressources"

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
    context.SetVar("fileOp",&FileTest{Debugcode : string(dat),Console : "",Backtrace : ""})

	
    window := component.CreateWindow(nil)
	

    window.Show()
    window.Wait()

    return nil
}

type FileTest struct {
    Debugcode   string
    Console string
    Backtrace string
}
/*
type outputGDB struct {
	pos int32
}

func (out *outputGDB) Write(p []byte) (n int, err error) {
	out.pos++
	fmt.Println(string(p))
	return len(p),nil
}
*/
//func (fileOp *FileTest) Debugrun() {  
	/*rescueStdout := os.Stdout
 	 r, w, _ := os.Pipe()
 	 os.Stdout = w*/

 
	
	//io.Copy(os.Stdout, debug)
	
	//fmt.Println("Hello, playground") 
	//debug.Exit()
  	/*w.Close()
	debug.Interrupt()
  	out, _ := ioutil.ReadAll(r)
  	os.Stdout = rescueStdout*/

 	/*out := &outputGDB{}
	start()
	io.Copy(out, debug)
	debug.Exit()*/
	



	// fmt.Printf("Captured: %s \n", out)
	//debug.Interrupt()
  	//fileOp.Console = fileOp.Console + "\n" + string(out)
	//qml.Changed(consoleOp, &fileOp.Console)
	//debug.Exit()



func (fileOp *FileTest) Debugrun() {  

     	start() 
	where()
	
	
	//Backtrace
	fileOp.Backtrace = backtrace()
	qml.Changed(fileOp, &fileOp.Backtrace)
	
	
	

	


}
func (fileOp *FileTest) Debugstep() {  
	
        step()
	 
	where()
	

	//Backtrace
	fileOp.Backtrace = backtrace()
	qml.Changed(fileOp, &fileOp.Backtrace)

}
func (fileOp *FileTest) Debugcontinue() {  
	
        continuee() 
	

	where()


	//Backtrace
	fileOp.Backtrace = backtrace()
	qml.Changed(fileOp, &fileOp.Backtrace)

	
}

func (fileOp *FileTest) Debugreverse() {  
        
	step_reverse()
	where()
	
	//Backtrace
	fileOp.Backtrace = backtrace()
	qml.Changed(fileOp, &fileOp.Backtrace)
}

func (fileOp *FileTest) Addbreakpoint(bp int) {  
        
	
	breake(bp)
	
}
func (fileOp *FileTest) Rmvbreakpoint(bp int){
	delete_break(bp)
	
}


func (fileOp *FileTest) Debugreversecontuinue(){
	continue_reverse()
	where()
	//Backtrace
	fileOp.Backtrace = backtrace()
	qml.Changed(fileOp, &fileOp.Backtrace)
}

/*func (fileOp *FileTest) Debugstop(){
	stop()
}*/

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



