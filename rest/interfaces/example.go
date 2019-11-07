package interfaces

type CarDatabaseInterface interface {
	AddCar(c Car)
	SearchCar(plate string) (*Car, error)
	UpdateCar(c Car) error
	DeleteCar(plate string) error
}

type MyInterface interface {
	MethodA()
	MethodB(value string) (string, error)
}

type CustomType struct {}

func (c CustomType) MethodA() {
	panic("implement me")
}

func (c CustomType) MethodB(value string) (string, error) {
	panic("implement me")
}

type Car struct {
	Plate string `json:"plate"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func (c *CarDatabase) AddCar(car Car) {
	panic("implement me")
}

func (c *CarDatabase) SearchCar(plate string) (*Car, error) {
	panic("implement me")
}

func (c *CarDatabase) UpdateCar(car Car) error {
	panic("implement me")
}

func (c *CarDatabase) DeleteCar(plate string) error {
	panic("implement me")
}

type CarDatabase struct {
	Cars []Car
}

type CarDatabaseWithName struct {
	Name string
	CarDatabase
}

type MySimpleInterface interface {
	SimpleMethod()
}

type MyComplexInterface interface {
	MySimpleInterface
	ComplexMethod()
}

type MyType struct {}

func (m MyType) SimpleMethod() {
	panic("implement me")
}

func (m MyType) ComplexMethod() {
	panic("implement me")
}

func Run() {

}
