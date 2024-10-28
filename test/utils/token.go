package utils

import (
	"encoding/json"
	"log"
	"crypto/rsa"
	"time"
	"encoding/base64"
	"math/big"

	"github.com/golang-jwt/jwt/v5"
)

var (
	d string = "jFhfsR-4UWSMbl-47HN5jG_mwQT0-fqrTPlqV5foQCsOmrPIHZC5LQEBCNo6_72nyNryqdzTvEesgYK0EY0vgFyderIn4ajOQZywSL1Pa4uu8HXmycpH50K40i-7RyVnMPeuoPuApM4xWrKg_9SmM4GgVxj84nCZo3y7jMIk76z8fBcF7_FiRHIim9bFAXx0b8BCRusySMjXwqcO6C8JZ8BqfEQruG5YgVQ293pGNzZKV-jYBYg0Q5hxfhtyT5ocIEXldDaFbQRwKxY8XRz0FtOoolYyNR7ujgGhZR5BOnn_lD4-ajpS6i6YrOnQH2V9z7H6l79wTSbQtB7dhvin8Q"
	e string = "AQAB"
	n string = "zURS2VvS9sc6QC6pMJqb41_A1z8c7lPF-a9zyFI4-7eHDS6yueBDFQUIkxkBbCLkpIf4TaZqa1LZx7iYRJxLbBlUxfg16HorVLD1SNt6EOrNIwJa8KHORS2hnNt7-XH5K1D4UdMGJ6GQD9YzcZeh8JXAfEeequ8kFyPspwQg_JWaMvoi3AlEUesbZ2RflktuS0wd6M50tRb-Kx0L60jrPP_O-z-Epwy41Yng-SUciecaAc1Ms1W5K_Jx6dJDdSvTwKnM9WdEpL7_7Z-G8vlpfm4M0LI8AiibxAcFLnYqfzgmwh5eOZ6aU3IraZaTHB5g0X-tXRNbl-O_jt7tbW6_vw"
)

func GetJwksString() string {
	set := map[string]interface{}{
		"kid": "0",
		"kty": "RSA",
		"e":   e,
		"n":   n,
	}
	outerMap := map[string]interface{}{
		"keys": []interface{}{set},
	}
	jsonData, _ := json.Marshal(outerMap)

	return string(jsonData)
}

func GenerateToken(role string) string {
	privKey := createPrivateKeyFromJWK(map[string]string{"n": n, "e": e, "d": d})
	header := map[string]interface{}{"typ": "JWT", "alg": "RS256", "kid": "0"}
	payload := map[string]interface{}{
		"iss":      "api-auth",
		"iat":      time.Now().Unix(),
		"exp":      time.Now().AddDate(1, 0, 0).Unix(),
		"roles":    []string{role},
		"username": "test",
		"location": []interface{}{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(payload))
	token.Header = header

	tokenString, err := token.SignedString(privKey)
	if err != nil {
		log.Fatal("Error signing token:", err)
	}

	return tokenString
}

func createPrivateKeyFromJWK(jwk map[string]string) *rsa.PrivateKey {
	nBytes, _ := base64.RawURLEncoding.DecodeString(jwk["n"])
	eBytes, _ := base64.RawURLEncoding.DecodeString(jwk["e"])
	dBytes, _ := base64.RawURLEncoding.DecodeString(jwk["d"])

	e := big.NewInt(0).SetBytes(eBytes).Int64()
	n := new(big.Int).SetBytes(nBytes)
	d := new(big.Int).SetBytes(dBytes)

	privKey := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: n,
			E: int(e),
		},
		D:      d,
		Primes: []*big.Int{big.NewInt(0), big.NewInt(0)}, // You'd need the actual primes (p and q) to complete this
	}

	return privKey
}
