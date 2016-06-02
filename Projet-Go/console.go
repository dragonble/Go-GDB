package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"path/filepath"
)

func main() {
<<<<<<< HEAD
	cmd := exec.Command("../Projet-Go","gcc -g essai.exe -o console")
=======
	dir, err := filepath.Abs(filepath.Dir(os.Args[1]))
   	 if err != nil {
            log.Fatal(err)
    	}
	cmd := exec.Command(dir,"gcc -g " + os.Args[1] + " -o progtest")
>>>>>>> f89a2d39986625a3c90f2ecaeb6024b589500376
	
	
	out, err2 := cmd.Output()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("compile", out)
}
