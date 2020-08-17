package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Port string

func Initialize() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	Port = os.Getenv("port")
}
