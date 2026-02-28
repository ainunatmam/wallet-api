package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	bootstrap "wallet-api/bootsrap"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {

	
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env:", err)
	}
	
	
	mysql, err := bootstrap.NewDatabase()
	if err != nil {
		log.Fatal("Failed to open DB:", err)
		panic(err)
	}

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		runMigration(mysql)
		return
	}

	app := fiber.New()
	err = bootstrap.NewBootstrap(app, mysql).Run()
	if err != nil {
		panic(err)
	}
}

func runMigration(db *sql.DB) {
	if len(os.Args) < 3 {
		log.Println("Usage:")
		fmt.Println("  migrate up")
		fmt.Println("  migrate down")
		fmt.Println("  migrate status")
		fmt.Println("  migrate create <name>")
		log.Fatal("Specify migrate command")
	}

	action := os.Args[2]
	goose.SetDialect("mysql")
	switch action {
	case "up":
		if err := goose.Up(db, "./database/migration"); err != nil {
			log.Fatal(err)
		}

	case "down":
		if err := goose.Down(db, "./database/migration"); err != nil {
			log.Fatal(err)
		}

	case "status":
		if err := goose.Status(db, "./database/migration"); err != nil {
			log.Fatal(err)
		}

	case "create":
		if len(os.Args) < 4 {
			log.Fatal("Specify migration name")
		}
		name := os.Args[3]

		if err := goose.Create(db, "./database/migration", name, "sql"); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal("Unknown migrate action")
	}
}
