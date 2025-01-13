package main

import (
	"fmt"
)

func main() {
	simpleBuilder := "simple"
	richBuilder := "rich"
	director := NewDirector()
	builder, _ := getBuilder(simpleBuilder)
	director.setBuilder(builder)
	simpleHouse := director.buildHouse()
	fmt.Printf("\nЭтажей: %d\nОкна: %s\nУнитаз: %s", simpleHouse.floor, simpleHouse.windowsType, simpleHouse.toilet)

	builder, _ = getBuilder(richBuilder)
	director.setBuilder(builder)
	richHouse := director.buildHouse()
	fmt.Printf("\nЭтажей: %d\nОкна: %s\nУнитаз: %s", richHouse.floor, richHouse.windowsType, richHouse.toilet)

}
