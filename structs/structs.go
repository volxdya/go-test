package structs

import "math"

type Point struct {
	X float64
	Y float64
}

type User struct {
	FirstName string
	LastName  string
	Password  string
}

type Developer struct {
	PL   string
	User User
}

type Round struct {
	Radius float64
}

// * - указатель, нужен чтобы обращался исключительно к экземляру структуры
func (p *Point) SetXY(x, y float64) {
	p.X = x
	p.Y = y
}

func (p Point) Square() float64 {
	return p.X * p.Y
}

func (r Round) Square() float64 { return math.Pi * math.Pow(r.Radius, 2) }

type SuperFigure struct {
	Round
	Point
}

type F interface {
	Square() float64
}
