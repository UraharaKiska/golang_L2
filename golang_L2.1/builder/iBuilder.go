package main

import "errors"

type HouseBuilder interface {
	setWindowType()
	setFloorsCount()
	setToilet()
	getHouse() House
}

func getBuilder(builderType string) (HouseBuilder, error) {
	if builderType == "rich" {
		return &PenthouseBuilder{}, nil
	} else if builderType == "simple" {
		return &SimpleHouseBuilder{}, nil
	}
	return nil, errors.New("Cant identify builder")
}
