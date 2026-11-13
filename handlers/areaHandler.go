package handlers

type Area struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (a Area) Contains(other Area) bool {
	return other.X >= a.X &&
		other.Y >= a.Y &&
		other.X+other.Width <= a.X+a.Width &&
		other.Y+other.Height <= a.Y+a.Height
}
func (a Area) IsIn(x, y float64) bool {
	return x >= a.X &&
		x <= a.X+a.Width &&
		y >= a.Y &&
		y <= a.Y+a.Height
}
