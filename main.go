package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vannguyen2606/poseidon-core/database"
	"github.com/vannguyen2606/poseidon-core/routes"
)

func main() {
	// read file .env
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	//connect database
	database.Connect(myEnv)
	// init router
	app := fiber.New()
	routes.Setup(app)
	// connect port
	fmt.Printf("Listening to port %v", myEnv["PORT"])
	app.Listen(fmt.Sprintf(":%v", myEnv["PORT"]))

}
