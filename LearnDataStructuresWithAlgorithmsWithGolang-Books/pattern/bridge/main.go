package main

import "fmt"

type IDrawShape interface {
	drawShape(x[5] float32, y[5] float32)
}

type IContour interface {
	drawContour(x[5] float32, y[5] float32)
	resizeByFactor(factor int)
}

type DrawShape struct {}

type DrawContour struct {
	x[5] float32
	y[5] float32
	shape DrawShape
	factor int
}

func (draw DrawShape) drawShape(x[5] float32, y[5] float32)  {
	fmt.Println("draw shape")
}

func (contour DrawContour) drawContour(x[5] float32, y[5] float32)  {
	fmt.Println("drawing contour")
	contour.shape.drawShape(contour.x, contour.y)
}

func (contour DrawContour) resizeByFactor(factor int)  {
	contour.factor = factor
}

func main() {
	var x = [5]float32{1,2,3,4,5}
	var y = [5]float32{1,2,3,4,5}
	var contour = DrawContour{x,y,DrawShape{}, 2}

	contour.drawContour(x,y)
	contour.resizeByFactor(2)
}
