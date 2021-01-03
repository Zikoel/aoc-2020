package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowEquality(t *testing.T) {
	assert := assert.New(t)

	row1 := []string{"#", ".", "L"}
	assert.True(rowEquality(row1, row1))

	row2 := []string{"#", "#", "L"}
	assert.False(rowEquality(row1, row2))
	assert.False(rowEquality(row2, row1))

	row3 := []string{"#", ".", "L", "."}
	assert.False(rowEquality(row1, row3))
	assert.False(rowEquality(row3, row1))
}

func TestRoomFromString(t *testing.T) {
	assert := assert.New(t)

	roomStr := "#..\n.#.\n..L"

	room := roomFromString(roomStr)
	// fmt.Printf("%v", room)

	assert.Len(room, 3)
	assert.Len(room[0], 3)
	assert.Len(room[1], 3)
	assert.Len(room[2], 3)
	assert.Equal("#", room[0][0])
	assert.Equal(".", room[0][1])
	assert.Equal(".", room[0][2])
	assert.Equal(".", room[1][0])
	assert.Equal("#", room[1][1])
	assert.Equal(".", room[1][2])
	assert.Equal(".", room[2][0])
	assert.Equal(".", room[2][1])
	assert.Equal("L", room[2][2])
}

func TestRoomToString(t *testing.T) {
	assert := assert.New(t)

	roomStr := "#..\n.#.\n..L"

	room := roomFromString(roomStr)
	result := roomToString(room)

	assert.Equal(roomStr, result)
}

func TestIsPositionOccupied(t *testing.T) {
	assert := assert.New(t)

	roomStr := `#..
.#.
..L`

	room := roomFromString(roomStr)

	assert.True(isPositionOccupied(room, position{0, 0}))
	assert.False(isPositionOccupied(room, position{0, 1}))
	assert.False(isPositionOccupied(room, position{0, 2}))
	assert.False(isPositionOccupied(room, position{1, 0}))
	assert.True(isPositionOccupied(room, position{1, 1}))
	assert.False(isPositionOccupied(room, position{2, 2}))

}

func TestOccupiedAround(t *testing.T) {
	assert := assert.New(t)

	roomStr := `#..
.#.
..L`

	room := roomFromString(roomStr)

	assert.Equal(1, occupiedAround(room, 0, 0))
	assert.Equal(1, occupiedAround(room, 1, 1))
	assert.Equal(2, occupiedAround(room, 0, 1))
}

func TestRoomUpdate(t *testing.T) {
	assert := assert.New(t)

	roomStr := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	expected := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

	room := roomFromString(roomStr)
	updated := roomUpdate(room)
	result := roomToString(updated)

	// fmt.Printf("%v", room)

	assert.Equal(expected, result)
}
