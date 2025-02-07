package goflappy

type Position struct {
	X, Y float64
}

type Rect struct {
	position      Position
	Width, Height float64
}

func (r Rect) Overlaps(other Rect) bool {
	return r.position.X < other.position.X+other.Width &&
		r.position.X+r.Width > other.position.X &&
		r.position.Y < other.position.Y+other.Height &&
		r.position.Y+r.Height > other.position.Y
}
