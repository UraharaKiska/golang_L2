// You can edit this code!
// Click here and start typing.
package abstract_factory

import (
	"errors"
	"fmt"
	"os"
)

type GuiApplication interface {
	showButton()
	openWindow()
	getInfo()
}

type GuiWindows struct {
	width  uint
	height uint
}

func (w *GuiWindows) showButton() {
	fmt.Println("Winwows button")
}

func (w *GuiWindows) openWindow() {
	fmt.Println("Open Windows window")
}
func (w *GuiWindows) getInfo() {
	fmt.Println("Hi, i am in Windows")
}

type GuiMacOs struct {
	width  uint
	height uint
}

func (w *GuiMacOs) showButton() {
	fmt.Println("MacOs button")
}

func (w *GuiMacOs) openWindow() {
	fmt.Println("Open MacOs window")
}

func (w *GuiMacOs) getInfo() {
	fmt.Println("Hi, i am in MacOs")
}

func initApp(osType string) (GuiApplication, error) {
	if osType == "Win" {
		return &GuiWindows{0, 0}, nil
	} else if osType == "Mac" {
		return &GuiMacOs{0, 0}, nil
	} else {
		return nil, errors.New("Unvailable OS type")
	}
}

func main() {
	typeOs := "Mac"
	Application, err := initApp(typeOs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Application.getInfo()
	Application.openWindow()
	Application.showButton()

}
