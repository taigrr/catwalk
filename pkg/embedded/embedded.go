// Package embedded provides access to all providers in a embedded manner.
// This basically means offline access to the providers.
package embedded

import (
	"github.com/taigrr/catwalk/internal/providers"
	"github.com/taigrr/catwalk/pkg/catwalk"
)

// GetAll returns all embedded providers.
func GetAll() []catwalk.Provider {
	return providers.GetAll()
}
