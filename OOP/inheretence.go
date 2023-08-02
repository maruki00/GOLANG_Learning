package main

import "fmt"

type Shape struct {
	x int
	y int
}

func (s *Shape) Area() int {
	return s.x * s.y
}

func (s *Shape) getX() int {
	return s.x
}

func (s *Shape) setX(x int) {
	s.x = x
}

func (Shape *Shape) setY(y int) {
	Shape.y = y
}

func (Shape *Shape) getY() int {
	return Shape.y
}

type Rectangle struct {
	Shape
}

type Square struct {
	Shape
}

func main() {
	r := Rectangle{
		Shape: Shape{
			x: 10,
			y: 20,
		},
	}
	fmt.Println(r.getX())
	r.setX(20)
	fmt.Println(r.getX())
	fmt.Println(r.getY())
	r.setY(30)
	fmt.Println(r.getX())
	fmt.Println(r.getY())

	s := Square{
		Shape{
			x: 10,
			y: 20,
		},
	}
	fmt.Println(s.getX())
	s.setX(20)
	fmt.Println(s.getX())
	fmt.Println(s.getY())
	s.setY(30)
	fmt.Println(s.getX())
	fmt.Println(s.getY())
	fmt.Println(s.Area())
	fmt.Println(s.getX())
	fmt.Println(s.getY())
	fmt.Println(s.Area())
}
