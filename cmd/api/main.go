package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/m00n3r-dev/forgecom-api/internal/config"
	"github.com/m00n3r-dev/forgecom-api/internal/database"
	"github.com/m00n3r-dev/forgecom-api/internal/user"
)

func main() {

	cnf, err := config.Load()
	if err != nil {
		log.Fatal("failed to load config \n", err)
	}

	db, err := database.Connect(cnf)
	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations \n", err)
	}

	// router
	r := chi.NewRouter()

	// dependencies
	userRepository := user.NewRepository(db.DB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	r.Post("/auth/register", userHandler.Register)

	fmt.Printf("Application running on PORT :%s\n", cnf.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cnf.Port), r)
}
