package providers

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	AuthDomain = os.Getenv("AUTH0_DOMAIN")
	AuthAudience = os.Getenv("AUTH0_AUDIENCE")
)

func Auth0TokenMiddleware() gin.HandlerFunc {
	// Obtain the issuer URL from the Auth0 domain.
	issuerURL, err := url.Parse("https://" + AuthDomain + "/")

	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	// Create a new caching provider with a 5 minute cache.
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	// Create a new JWT validator.
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{AuthAudience},
		validator.WithAllowedClockSkew(time.Minute),
	)

	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	// Create an error handler.
	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{ "error": "Unauthorized" }`))
	}

	// Create a new JWT middleware.
	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	// Return the middleware as a gin.HandlerFunc.
	return func(c *gin.Context) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		})

		middleware.CheckJWT(handler)
	}
}
