package day12

import (
	"fmt"
	"math"
	"strconv"
)

func goForward(p position, val int32) position {

	newP := p

	switch p.angle {
	case 0:
		newP.n += val
	case 90:
		newP.e += val
	case 180:
		newP.s += val
	case 270:
		newP.w += val
	}

	return newP
}

func nextPosition(actual position, instruction string) position {
	action := instruction[:1]
	val, _ := strconv.ParseInt(instruction[1:], 10, 32)
	val32 := int32(val)

	var newPosition position = actual

	switch action {
	case "N":
		newPosition.n = actual.n + val32
	case "S":
		newPosition.s = actual.s + val32
	case "E":
		newPosition.e = actual.e + val32
	case "W":
		newPosition.w = actual.w + val32
	case "L":
		newPosition.angle = subAngle(actual.angle, val32)
	case "R":
		newPosition.angle = sumAngle(actual.angle, val32)
	case "F":
		newPosition = goForward(newPosition, val32)
	}

	return newPosition
}

func Run(data []byte) {
	instructions := toInstructions(data)
	position := position{
		n:     0,
		e:     0,
		w:     0,
		s:     0,
		angle: 90,
	}
	for _, instruction := range instructions {
		position = nextPosition(position, instruction)
		fmt.Println("%V", position)
	}

	manhattanDistance := math.Abs(float64(position.e-position.w)) + math.Abs(float64(position.n-position.s))

	fmt.Printf("Manhattan distance: %f", manhattanDistance)
}
