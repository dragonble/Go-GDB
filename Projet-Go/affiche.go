package main

import (
    "fmt"
    "gopkg.in/qml.v1"
    "github.com/cyrus-and/gdb"
    "io"
    "os"
)

func main() {

    





    if err := qml.Run(run); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {
    engine := qml.NewEngine()

    component, err := engine.LoadFile("texte.qml")
    if err != nil {
        return err
    }

    context := engine.Context()
    context.SetVar("ctrl", &Control{Name: "Enter your name"})

    window := component.CreateWindow(nil)

    window.Show()
    window.Wait()

    return nil
}

type Control struct {
    Name    string
    Message string
}

func (ctrl *Control) Hello() {
    go func() {
        ctrl.Message = "Hello, " + ctrl.Name
        qml.Changed(ctrl, &ctrl.Message)
    }()
}
