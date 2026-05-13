// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestFeaturesRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features", "retrieve",
			"--type", "type",
			"--id", "0",
		)
	})
}

func TestFeaturesBatch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features", "batch",
			"--element", "{id: 21154906, type: node}",
			"--element", "{id: 4589123, type: way}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(featuresBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features", "batch",
			"--element.id", "21154906",
			"--element.type", "node",
			"--element.id", "4589123",
			"--element.type", "way",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"elements:\n" +
			"  - id: 21154906\n" +
			"    type: node\n" +
			"  - id: 4589123\n" +
			"    type: way\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features", "batch",
		)
	})
}

func TestFeaturesQuery(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"features", "query",
			"--cursor", "cursor",
			"--format", "format",
			"--h3", "h3",
			"--limit", "0",
			"--type", "type",
			"--around", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--contains", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--crosses", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--intersects", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--not-contains", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--not-intersects", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--not-within", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--radius", "500",
			"--touches", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--within", "{coordinates: [2.3522, 48.8566], type: Point}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"around:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"contains:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"crosses:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"intersects:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"not_contains:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"not_intersects:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"not_within:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"radius: 500\n" +
			"touches:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"within:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"features", "query",
			"--cursor", "cursor",
			"--format", "format",
			"--h3", "h3",
			"--limit", "0",
			"--type", "type",
		)
	})
}
