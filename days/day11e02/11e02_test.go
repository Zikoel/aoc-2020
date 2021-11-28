package day11e02

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

	roomStr1 := `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`

	room1 := roomFromString(roomStr1)
	assert.Equal(8, occupiedAlongDirections(room1, 4, 3))

	roomStr2 := `#...
.L..
.##.
.#.#`

	room2 := roomFromString(roomStr2)
	assert.Equal(3, occupiedAlongDirections(room2, 1, 1))

	roomStr3 := `L..#
..#.
.#..
...#`

	room3 := roomFromString(roomStr3)
	assert.Equal(2, occupiedAlongDirections(room3, 0, 0))

	roomStr4 := `..#.
#...
.#..
..L.`

	room4 := roomFromString(roomStr4)
	assert.Equal(2, occupiedAlongDirections(room4, 3, 2))

	roomStr5 := `#.#.
.L..
#LLL
..#.`

	room5 := roomFromString(roomStr5)
	assert.Equal(2, occupiedAlongDirections(room5, 2, 2))

}

func TestRoomUpdate(t *testing.T) {
	assert := assert.New(t)

	roomStr1 := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	expected1 := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

	room := roomFromString(roomStr1)
	updated := roomUpdate(room)
	result := roomToString(updated)
	assert.Equal(expected1, result)

	roomStr2 := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

	expected2 := `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`

	room = roomFromString(roomStr2)
	updated = roomUpdate(room)
	result = roomToString(updated)
	assert.Equal(expected2, result)

	roomStr3 := `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`

	expected3 := `#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`

	room = roomFromString(roomStr3)
	updated = roomUpdate(room)
	result = roomToString(updated)
	assert.Equal(expected3, result)

}
