package main

import (
	"log"
	"os"

	"github.com/CS559-CSD-IITBH/search-service/routes"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Internal server error: Unable to load the env file")
	}

	store := sessions.NewFilesystemStore("sessions/", []byte("secret-key"))

	store.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	r := routes.SetupRouter(store)
	r.Run(":" + os.Getenv("PORT"))
}
