package day11e02

import (
	"fmt"
	"strings"
)

type room = [][]string

type position struct {
	row    int
	column int
}

func isPositionIsideRoom(room room, p position) bool {

	if p.row < 0 || p.row >= len(room) {
		return false
	}

	if p.column < 0 || p.column >= len(room[p.row]) {
		return false
	}

	return true
}

func isPositionOccupied(r room, p position) bool {

	if !isPositionIsideRoom(r, p) {
		return false
	}

	return r[p.row][p.column] == "#"
}

func isPositionFree(r room, p position) bool {

	if !isPositionIsideRoom(r, p) {
		return false
	}

	return r[p.row][p.column] == "L"
}

func occupiedAlongDirections(room room, r, c int) int {
	directionFormula := []position{
		{-1, +0}, // N
		{-1, +1}, // NE
		{+0, +1}, // E
		{+1, +1}, // SE
		{+1, +0}, // S
		{+1, -1}, // SO
		{+0, -1}, // O
		{-1, -1}, // NO
	}

	occupied := 0
	for _, formula := range directionFormula {
		p := position{r, c}
		for isPositionIsideRoom(room, p) {
			if isPositionOccupied(room, p) && (p.row != r || p.column != c) {
				occupied++
				break
			}
			if isPositionFree(room, p) && (p.row != r || p.column != c) {
				break
			}
			p.column = p.column + formula.column
			p.row = p.row + formula.row
		}
	}

	return occupied
}

func roomUpdate(rr room) room {
	updated := room{}
	for idr, r := range rr {
		row := []string{}
		for idc, c := range r {
			n := occupiedAlongDirections(rr, idr, idc)
			if c == "L" && n == 0 {
				row = append(row, "#")
			} else if c == "#" && n >= 5 {
				row = append(row, "L")
			} else {
				row = append(row, c)
			}
		}
		updated = append(updated, row)
	}
	return updated
}

func roomFromString(data string) room {
	rows := strings.Split(data, "\n")

	room := room{}
	for _, r := range rows {
		newRow := []string{}
		columns := strings.Split(r, "")
		newRow = append(newRow, columns...)
		room = append(room, newRow)
	}

	return room
}

func rowEquality(row1, row2 []string) bool {

	if len(row1) != len(row2) {
		return false
	}

	result := true
	for idx, c := range row1 {
		if row2[idx] != c {
			result = false
		}
	}

	return result
}

func roomsEquality(room1, room2 room) bool {

	if len(room1) != len(room2) {
		return false
	}

	result := true

	for idx, r := range room1 {
		if !rowEquality(r, room2[idx]) {
			result = false
		}
	}

	return result
}

func roomToString(room room) string {
	result := ""
	for _, r := range room {
		for _, c := range r {
			result += fmt.Sprintf("%s", c)
		}
		result += fmt.Sprint("\n")
	}

	return strings.TrimSuffix(result, "\n")
}

func countOccupuedSit(room room) int {
	occupiedSit := 0
	for _, r := range room {
		for _, c := range r {
			if c == "#" {
				occupiedSit++
			}
		}
	}
	return occupiedSit
}

// Run run
func Run(data []byte) {
	room := roomFromString(string(data))

	// Start
	maxIterations := 1000
	roomIsSame := false
	precRoomIteration := room
	for maxIterations > 0 && roomIsSame == false {
		newRoom := roomUpdate(precRoomIteration)
		if roomsEquality(newRoom, precRoomIteration) == true {
			fmt.Println("rooms are equals")
			roomIsSame = true
		}
		precRoomIteration = newRoom
		maxIterations--
	}

	fmt.Printf("%s\n", roomToString(precRoomIteration))
	fmt.Printf("Stability after %d\n", 1000-maxIterations)
	fmt.Printf("Occupied sits: %d", countOccupuedSit(precRoomIteration))
}
