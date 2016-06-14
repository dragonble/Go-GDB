package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {
    dir, err := filepath.Abs(filepath.Dir(os.Args[2]))
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println(dir)
}
