package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type  compDash struct{

}

func main(){

	// fmt
	fmt.Println("Hello World")

	// app is an instance of Fiber
	//In Fiber, a Ctx (context) object is used to manage the state and information 
    //related to the current HTTP request.    
	app := fiber.New()

	// GET route heakthcheck route
	app.Get("/healthcheck", func(c *fiber.Ctx) error { return c.SendString("OK!")})
    //.SendString("") is a built in Fiber method to return a reponse as a string

	// Start Server  app listener localhost 3000
	// log fatal => logs the error and terminate the program
	log.Fatal(app.Listen(":3000"))


	app.Get()
	app.Post()
	app.Put()
	app.Delete()
}
