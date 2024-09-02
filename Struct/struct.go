package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func Perimenter(reactanfle Rectangle) float64 {
	return 2 * (reactanfle.Height + reactanfle.Width)
}

// (c Circle) <-これはメソッドレシーバー
// 慣例としてレシーバー変数の頭文字を変数名とする
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
