package structs

// Perimeter returns the perimeter of the rectangle.
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of the rectangle.
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
