package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("../Projet-Go","gcc -g essai.exe -o console")
	
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("compile", out)
}
