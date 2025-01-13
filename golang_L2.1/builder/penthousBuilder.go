package main

type PenthouseBuilder struct {
	windowType  string
	floorsCount int
	toilet      string
}

func (b *PenthouseBuilder) setWindowType() {
	b.windowType = "Панорамные окна"
}

func (b *PenthouseBuilder) setFloorsCount() {
	b.floorsCount = 3
}

func (b *PenthouseBuilder) setToilet() {
	b.toilet = "Японский зололтой унитаз"
}

func (b *PenthouseBuilder) getHouse() House {
	return House{
		floor:       b.floorsCount,
		windowsType: b.windowType,
		toilet:      b.toilet,
	}
}
