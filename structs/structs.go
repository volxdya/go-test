package structs

type Point struct {
	X float64
	Y float64
}

func (p Point) Square() float64 {
	return p.X * p.Y
}

// * - указатель, нужен чтобы обращался исключительно к экземляру структуры
func (p *Point) SetXY(x, y float64) {
	p.X = x
	p.Y = y
}
