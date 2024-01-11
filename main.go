package main

import (
    "os"
	"log"

	"github.com/joho/godotenv"
    "github.com/CS559-CSD-IITBH/search-service/routes"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatalln("Internal server error: Unable to load the env file")
	}

	r := routes.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
