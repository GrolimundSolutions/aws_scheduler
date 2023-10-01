package schedulermain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	// 3. Test missing or incorrectly formatted config file
	t.Run("missing or incorrect config file", func(t *testing.T) {
		_, err := LoadConfig("./invalidpath")
		assert.NotNil(t, err)
	})
}
