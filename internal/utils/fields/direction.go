package fields

type Direction int

const (
	DIRECTION_NONE  Direction = 0
	DIRECTION_EAST  Direction = 1
	DIRECTION_NORTH Direction = 2
	DIRECTION_WEST  Direction = 3
	DIRECTION_SOUTH Direction = 4
)

func (d Direction) String() string {
	switch d {
	case DIRECTION_EAST:
		return "E"
	case DIRECTION_NORTH:
		return "N"
	case DIRECTION_WEST:
		return "W"
	case DIRECTION_SOUTH:
		return "S"
	}
	return ""
}
