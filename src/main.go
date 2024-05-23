package main

import (
	events_mod "epsa.upv.es/postin_backend/src/modules/events_mod"
	profile_mod "epsa.upv.es/postin_backend/src/modules/profile_mod"
	proposes_mod "epsa.upv.es/postin_backend/src/modules/proposes_mod"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	route := gin.Default()

	// Initialize Middlewares
	route.Use(gin.Logger())
	route.Use(providers.Auth0TokenMiddleware())

	// Initialize Providers
	providers.InitDatabase()
	providers.InitCommandBus()
	providers.InitStorageClient()

	// Initialize Modules
	profile_mod.ProfileModule()
	proposes_mod.ProposesModule()
	events_mod.EventsModule()

	// Run the server
	route.Run(":8080")
}
