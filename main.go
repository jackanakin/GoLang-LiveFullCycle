package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	"database/sql"
)

func httpServer(){
	println("Hello world!")
	http.ListenAndServe(":8080", nil)
}

type Car struct {
	Name string `json:"name"`
	Model string `json:"model"`
	Price float64 `json:"price"`
}

func (c Car) Andar(){
	fmt.Println("O carro", c.Name, "está andando")
}

var cars []Car

func generateCars() {
	cars = append(cars, Car{"F2000", "Ferrari", 100})
	cars = append(cars, Car{"Aventador", "Lamborgini", 200})
	cars = append(cars, Car{"Cayene", "Porsche", 300})
}

func main() {
	//httpServer()

	nome := "Jardel"
	fmt.Println(nome)

	result, err := somar(1,2)
	fmt.Println(result, err)

	carro := Car{"Gol", "VW", 12}
	fmt.Println(carro.Name)

	carro.Andar()

	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, model, price) VALUES ($1, $2, $3")
	if err != nil {
		return nil
	}

	_, err = stmt.Exec(car.Name, car.Model, car.Price)
	if err != nil {
		return err
	}

	return nil
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func somar(a int, b int) (int, error) {
	if a+b > 10 {
		return 0, fmt.Errorf("soma maior que 10")
	}

	return a+b, nil
}