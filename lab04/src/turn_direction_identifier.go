package src

const (
	Left     string = "left"
	Right    string = "right"
	Straight string = "straight"
)

func GetTurnDirection(p1, p2, p3 Point) string {
	det := (p2.X-p1.X)*(p3.Y-p2.Y) - (p2.Y-p1.Y)*(p3.X-p2.X)

	if det > 0 {
		return Left
	} else if det < 0 {
		return Right
	} else {
		return Straight
	}
}
