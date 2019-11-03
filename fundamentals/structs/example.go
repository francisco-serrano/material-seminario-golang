package structs

type Car struct {
	Brand    string
	Model    string
	Speed    int32
	maxSpeed int32
}

func NewCar(brand, model string, maxSpeed int32) Car {
	return Car{
		Brand:    brand,
		Model:    model,
		Speed:    0,
		maxSpeed: maxSpeed,
	}
}

func (c *Car) Accelerate() {
	if c.Speed <= (c.maxSpeed - 10) {
		c.Speed += 10
	}
}

func (c *Car) Brake() {
	if c.Speed >= 10 {
		c.Speed -= 10
	}
}
