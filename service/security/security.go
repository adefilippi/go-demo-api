package security

import (
	"encoding/json"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
	"time"
	"example/web-service-gin/service/env"
	"example/web-service-gin/service/utils"
)

var (
	jwksURL string
	issuer  string // Update this to the expected issuer
	ApiKeys = strings.Split(env.GetEnvVariable("API_KEYS_WHITE_LIST"), ",")
)

// Function to fetch the JWKS keys from the remote server
func CheckApiKey(apiKey string) (bool, string) {
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
	jwksURL = env.GetEnvVariable("AUTH_SERVER_KEYSET_URL")
	issuer = env.GetEnvVariable("AUTH_SERVER_NAME")

	if bearer == "" {
		return false, "Header Authorization is missing"
	} else if jwksURL == "" {
		return false, "Auth server env var is missing"
	} else if issuer == "" {
		return false, "Auth server env var is missing"
	}

	splitToken := strings.Split(bearer, "Bearer ")
	if len(splitToken) != 2 {
		log.Fatalf("Invalid Bearer value")
		return false, "Invalid Bearer value"
	}
	tokenString := splitToken[1]
	if tokenString == "" {
		log.Fatalf("Invalid Bearer value")
		return false, "Invalid Bearer value"
	}

	jwks, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		log.Fatalf("Failed to create JWK Set from resource at the given URL.\nError: %s", err)
		return false, "Failed to create JWK Set from resource at the given URL."
	}

	jwksJSON := json.RawMessage(`{"keys":[{"e":"AQAB","kid":"0","kty":"RSA","n":"xbMECG1-JyQ5iY-qG24EX3AbU1wNE2CDqU9YhuKZ7CTYyM8sjorFDv5DvfEL6f_8Eegt7fUXcBIzsLGp7VTumovGrTWcwgS5DD1tj86M8R1-Ob6qB4bLnifR5GkoYyic1snr9e2EDdtN2yvw6jIg0S_B95elhtDJeJW6J2iNrTePbq-d59ezpQJ0MPFTYsaeYXCrITfhH1AjuRnLqgvrZ1sjackj7SS6Nw9x8qMELMEY1F4BqsMiAwASFhzLnYW-toLrcrTv8UNiFY8I-5lIoS_FQf5o-J6lPoZTxWKdejvk-f9mZ0DBG6kLxlcwkFKJbrEg9zNMC-eauwDJvMjNgw"}]}`)

	// Create the keyfunc.Keyfunc.
	jwks, err = keyfunc.NewJWKSetJSON(jwksJSON)
	if err != nil {
		log.Fatalf("Failed to create JWK Set from resource at the given URL.\nError: %s", err)
		return false, "Failed to create JWK Set from resource at the given URL."
	}

	token, err := jwt.Parse(tokenString,
		jwks.Keyfunc,
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(issuer))
	if err != nil || !token.Valid {
		log.Fatalf(issuer, err)
		return false, "Failed to parse token"
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Fatalf("Invalid claims")
		return false, "Invalid claims"
	}

	// Compare the "exp" claim to the current time
	expClaim, err := claims.GetExpirationTime()
	if err != nil {
		log.Fatalf("Failed to get exp. Error: %s", err)
		return false, "Failed to get exp"
	}

	if expClaim.Unix() < time.Now().Unix() {
		log.Fatalf("Invalid token, token is expired")
		return false, "Invalid token, token is expired"
	}

	return true, "ok"
}
