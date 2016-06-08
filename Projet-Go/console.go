package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"path/filepath"
)

func main() {

	cmd := exec.Command("../Projet-Go","gcc -g essai.exe -o console")

	dir, err := filepath.Abs(filepath.Dir(os.Args[1]))
   	 if err != nil {
            log.Fatal(err)
    	}
	cmd := exec.Command(dir,"gcc -g " + os.Args[1] + " -o progtest")

	
	
	out, err2 := cmd.Output()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("compile", out)
}
