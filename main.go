package main

import (
	"fga-final-project/config"
	"fga-final-project/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {	
	defer config.CloseDB()
	_ = godotenv.Load(".env")

	router.StartApp().Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}