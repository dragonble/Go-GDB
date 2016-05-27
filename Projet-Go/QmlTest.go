package main

import(
 	"gopkg.in/qml.v1"
	"fmt"
)

func main() {

        err := qml.Run(run)
	
        fmt.Println(err)
	
}

func run() error {
        engine := qml.NewEngine()
	
        component, err := engine.LoadFile("rectangle.qml")
	
        if err != nil {
                return err
        }
        win := component.CreateWindow(nil)
	win.On("visibleChanged", func(visible bool) {
        	if (visible) {
                	fmt.Println("Width:", win.Int("width"))
        	}
	})
        win.Show()

        win.Wait()
        return nil

}

