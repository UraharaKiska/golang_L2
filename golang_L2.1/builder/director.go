package main

type Director struct {
	builder HouseBuilder
}

func (d *Director) setBuilder(builder HouseBuilder) {
	d.builder = builder
}

func (d *Director) buildHouse() House {
	d.builder.setFloorsCount()
	d.builder.setWindowType()
	d.builder.setToilet()
	return d.builder.getHouse()
}

func NewDirector() Director {
	return Director{}
}
