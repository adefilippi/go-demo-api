package security

import (
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/MicahParks/keyfunc/v3"

	"example/web-service-gin/service/env"
	"example/web-service-gin/service/utils"
)

var (
	JwksURL string
	issuer  string // Update this to the expected issuer
	ApiKeys = make([]string, 0)
)

// Function to fetch the JWKS keys from the remote server
func CheckApiKey(apiKey string) (bool, string) {
	ApiKeys = strings.Split(env.GetEnvVariable("API_KEYS_WHITE_LIST"), ",")
	if apiKey == "" || ApiKeys == nil || len(ApiKeys) == 0 {
		// Check if bearer
		return false, "Authorization not present in header"
	}

	if !utils.Contains(ApiKeys, apiKey) {
		return false, "Invalid key provided"
	}

	return true, "ok"
}

func CheckBearer(bearer string) (bool, string) {
	JwksURL = env.GetEnvVariable("AUTH_SERVER_KEYSET_URL")
	issuer = env.GetEnvVariable("AUTH_SERVER_NAME")

	if bearer == "" {
		return false, "Header Authorization is missing"
	} else if JwksURL == "" {
		return false, "Auth server env var is missing"
	} else if issuer == "" {
		return false, "Auth server env var is missing"
	}

	splitToken := strings.Split(bearer, "Bearer ")
	if len(splitToken) != 2 {
		log.Println("Invalid Bearer value")
		return false, "Invalid Bearer value"
	}
	tokenString := splitToken[1]
	if tokenString == "" {
		log.Println("Invalid Bearer value")
		return false, "Invalid Bearer value"
	}
	jwks, err := keyfunc.NewDefault([]string{JwksURL})
	if err != nil {
		log.Println("Failed to create JWK Set from resource at the given URL.\nError: %s", err)
		return false, "Failed to create JWK Set from resource at the given URL."
	}

	token, err := jwt.Parse(tokenString,
		jwks.Keyfunc,
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(issuer))
	if err != nil {
		log.Println("%s", err)
		return false, "Failed to parse token"
	}

	if !token.Valid || token == nil {
		log.Println("Invalid token")
		return false, "Invalid token"
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Invalid claims")
		return false, "Invalid claims"
	}

	// Compare the "exp" claim to the current time
	expClaim, err := claims.GetExpirationTime()
	if err != nil {
		log.Println("Failed to get exp. Error: %s", err)
		return false, "Failed to get exp"
	}

	if expClaim.Unix() < time.Now().Unix() {
		log.Println("Invalid token, token is expired")
		return false, "Invalid token, token is expired"
	}

	return true, "ok"
}
