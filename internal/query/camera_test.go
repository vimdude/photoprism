package query

import (
	"testing"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/stretchr/testify/assert"
)

// Example for using database fixtures defined in assets/resources/examples/fixtures.sql
func TestCamera_FirstOrCreate(t *testing.T) {
	t.Run("iphone-se", func(t *testing.T) {
		camera := entity.NewCamera("iPhone SE", "Apple")
		c := config.TestConfig()
		camera.FirstOrCreate(c.Db())
		assert.GreaterOrEqual(t, camera.ID, uint(1))
	})
}
