package main

import (
	"log"

	"arifudin-golang-learn/internal/config"
	"arifudin-golang-learn/internal/db"
	"arifudin-golang-learn/internal/handlers"
	kafkapkg "arifudin-golang-learn/internal/kafka"
	"arifudin-golang-learn/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	dbConn, err := db.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	prod, err := kafkapkg.NewProducer(cfg.KafkaBootstrap, cfg.KafkaTopic)
	if err != nil {
		log.Fatal(err)
	}
	defer prod.Close()

	repo := repository.NewUserRepository(dbConn)
	userHandler := handlers.NewUserHandler(repo, prod)

	app := fiber.New()

	api := app.Group("/api")
	u := api.Group("/users")
	u.Post("/", userHandler.Create)
	u.Get("/", userHandler.List)
	u.Get("/:id", userHandler.Get)
	u.Put("/:id", userHandler.Update)
	u.Delete("/:id", userHandler.Delete)

	log.Fatal(app.Listen(":" + cfg.Port))
}
