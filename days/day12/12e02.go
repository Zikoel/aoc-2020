package day12

import (
	"fmt"
	"math"
	"strconv"
)

// cartesian point
type cp struct {
	x int32
	y int32
}

func rotatePoint(p cp, angle int32, orientation string) cp {

	newP := p
	var rotateFn func(cp) cp

	if orientation == "CW" {
		rotateFn = func(p cp) cp { return cp{x: p.y, y: -p.x} }
	} else {
		rotateFn = func(p cp) cp { return cp{x: -p.y, y: p.x} }
	}

	switch angle {
	case 90:
		newP = rotateFn(newP)
	case 180:
		newP = rotateFn(newP)
		newP = rotateFn(newP)
	case 270:
		newP = rotateFn(newP)
		newP = rotateFn(newP)
		newP = rotateFn(newP)
	}
	return newP

}

func applyInstruction(boat, way cp, instruction string) (cp, cp) {
	action := instruction[:1]
	val64, _ := strconv.ParseInt(instruction[1:], 10, 32)
	val := int32(val64)

	newBoat := boat
	newWay := way

	switch action {
	case "N":
		newWay.y += val
	case "S":
		newWay.y -= val
	case "E":
		newWay.x += val
	case "W":
		newWay.x -= val
	case "L":
		newWay = rotatePoint(way, val, "CCW")
	case "R":
		newWay = rotatePoint(way, val, "CW")
	case "F":
		newBoat.x += way.x * val
		newBoat.y += way.y * val
	}

	return newBoat, newWay
}

func Run2(data []byte) {
	instructions := toInstructions(data)
	boat := cp{
		x: 0,
		y: 0,
	}
	way := cp{
		x: 10,
		y: 1,
	}

	for _, instruction := range instructions {
		boat, way = applyInstruction(boat, way, instruction)
		fmt.Printf("%v  | %v\n", boat, way)
	}

	manhattanDistance := math.Abs(float64(boat.x)) + math.Abs(float64(boat.y))

	fmt.Printf("Manhattan distance: %f", manhattanDistance)
}
