package main

type SimpleHouseBuilder struct {
	windowType  string
	floorsCount int
	toilet      string
}

func (b *SimpleHouseBuilder) setWindowType() {
	b.windowType = "Пластиковые окна"
}

func (b *SimpleHouseBuilder) setFloorsCount() {
	b.floorsCount = 1
}

func (b *SimpleHouseBuilder) setToilet() {
	b.toilet = "Белый керамический унитаз"
}

func (b *SimpleHouseBuilder) getHouse() House {
	return House{
		floor:       b.floorsCount,
		windowsType: b.windowType,
		toilet:      b.toilet,
	}
}
