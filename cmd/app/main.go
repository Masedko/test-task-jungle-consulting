package main

import (
	"log"
	"test-task-jungle-consulting/configuration"
	"test-task-jungle-consulting/internal/api/app"
)

// @title						Test task jungle consulting
// @version					1.0
// @description				API for Test task jungle consulting
// @securityDefinitions.apikey	JWT
// @in							header
// @name						Authorization
// @BasePath					/
func main() {
	err := configuration.LoadConfig(".env")
	if err != nil {
		log.Fatalln("Failed to load environment variables!", err.Error())
	}
	app.Run(&configuration.EnvConfig)
}
