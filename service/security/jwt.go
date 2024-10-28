package security

import (
	"github.com/MicahParks/keyfunc/v3"
)

// JwksProvider est une interface pour fournir le JWKSet
type JwksProvider interface {
	NewJwks(string) (keyfunc.Keyfunc, error)
}

// RealJwksProvider implémente l'interface JwksProvider en utilisant keyfunc
type ApiJwksProvider struct{}

func (p *ApiJwksProvider) NewJwks(url string) (keyfunc.Keyfunc, error) {
	return keyfunc.NewDefault([]string{url})
}
