package entity

import (
	"fmt"
	"sync"
)

// typeRegistry is a global map for type registration
var (
	typeRegistry = make(map[string]func() interface{})
	mu           sync.RWMutex
)

// RegisterType allows dynamic registration of types
func RegisterType(name string, factory func() interface{}) {
	mu.Lock()
	defer mu.Unlock()
	typeRegistry[name] = factory
}

// CreateStructFromString dynamically creates an instance by its name
func CreateStructFromString(typeName string) (interface{}, error) {
	mu.RLock()
	defer mu.RUnlock()

	factory, exists := typeRegistry[typeName]
	if !exists {
		return nil, fmt.Errorf("unknown type: %s", typeName)
	}

	return factory(), nil
}
