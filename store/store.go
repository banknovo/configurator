package store

import "github.com/banknovo/configurator/config"

// Store represents storage of parameters
type Store interface {
	// Fetches all the values which start with given path
	FetchAll(path string) ([]*config.Config, error)
}
