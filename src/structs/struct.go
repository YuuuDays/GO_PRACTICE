package structs

// 構造体
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of the rectangle.
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of the rectangle.
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
