package server

import (
	"rgb/internal/conf"
	"rgb/internal/database"
	"rgb/internal/store"
)

const InternalServerError = "Something went wrong!"

func Start(cfg conf.Config) {
	JwtSetup(cfg)
	store.SetDBConnection(database.NewDBOptions(cfg))

	router := setRouter()

	// Start listening and serving requests
	router.Run(":" + cfg.Port)
}
