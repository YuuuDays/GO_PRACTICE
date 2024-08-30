package main

type Reactanfle struct {
	Width  float64
	Height float64
}

func Perimenter(reactanfle Reactanfle) float64 {
	return 2 * (reactanfle.Height + reactanfle.Width)
}

func Area(x, y float64) float64 {
	return x * y
}
