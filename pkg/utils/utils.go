package utils

// BinarySplit 05
func BinarySplit(code string, space []int) int {
	head := string(code[0])

	// fmt.Printf("code %s space: %v\n", code, space)

	if len(space) == 2 {
		if head == "F" || head == "L" {
			return space[0]
		}
		return space[1]
	}

	half := len(space) / 2

	if head == "B" || head == "R" {
		return BinarySplit(code[1:], space[half:])
	}
	return BinarySplit(code[1:], space[:half])
}
