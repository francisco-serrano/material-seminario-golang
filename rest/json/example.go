package json

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SampleStruct struct {
	StatusCode   int            `json:"status_code"`
	Message      string         `json:"message"`
	NestedStruct *AnotherStruct `json:"nested_struct"`
}

type AnotherStruct struct {
	Message string `json:"message"`
}

type Car struct {
	Plate      string `json:"plate"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	EngineType string `json:"engine_type,omitempty"`
	Year       int    `json:"year"`
}

type CarDatabase struct {
	Cars []Car
}

func (c *CarDatabase) SearchCar(plate string) (*Car, error) {
	for _, car := range c.Cars {
		if car.Plate == plate {
			return &car, nil
		}
	}

	return nil, errors.New("car not found")
}

func Run() {
	database := CarDatabase{}

	car, err := database.SearchCar("ABC123")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(car)

	myCar := Car{
		Brand: "Volkswagen",
		Model: "Golf",
		Year:  2015,
	}

	marshalledCar, err := json.Marshal(myCar)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(myCar)
	fmt.Println(marshalledCar)
	fmt.Println(string(marshalledCar))

	var unmarshalledCar Car
	err = json.Unmarshal(marshalledCar, &unmarshalledCar)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(unmarshalledCar)
}
