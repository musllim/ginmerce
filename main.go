package main

import (
	"log"

	"github.com/lpernett/godotenv"
	"github.com/musllim/ginmerce/inits"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	inits.ConectDb()
	inits.Migrate()
}
func main() {

}
