package day12

import "strings"

type position struct {
	n     int32
	e     int32
	w     int32
	s     int32
	angle int32
}

func toInstructions(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func sumAngle(a, b int32) int32 {
	return (a + b) % 360
}

func subAngle(a, b int32) int32 {

	result := a - b

	for result < 0 {
		result += 360
	}

	return result
}
