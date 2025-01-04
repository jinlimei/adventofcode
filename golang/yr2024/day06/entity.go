package day06

type entity uint8

const (
	entityNothing entity = iota
	entityObstacle
	entityTempObstacle
	entityWalkPath
	entityGuard
)

func (co entity) String() string {
	switch co {
	case entityNothing:
		return "."
	case entityObstacle:
		return "#"
	case entityWalkPath:
		return "X"
	case entityGuard:
		return "g"
	case entityTempObstacle:
		return "O"
	default:
		panic("invalid entity")
	}
}
