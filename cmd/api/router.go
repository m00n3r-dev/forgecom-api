package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/m00n3r-dev/forgecom-api/internal/auth"
	"github.com/m00n3r-dev/forgecom-api/internal/config"
	"github.com/m00n3r-dev/forgecom-api/internal/database"
	"github.com/m00n3r-dev/forgecom-api/internal/middleware"
	"github.com/m00n3r-dev/forgecom-api/internal/user"
)

func SetupRoutes(config *config.Config, db *database.DB) *chi.Mux {

	jwtService := auth.NewJwtService(config.JwtSecret)
	refreshTokenRepo := auth.NewRefreshTokenRepository(db.DB)

	// router
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)

	// dependencies
	userRepository := user.NewRepository(db.DB)
	userService := user.NewService(userRepository, jwtService, refreshTokenRepo)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(r)

	return r
}
