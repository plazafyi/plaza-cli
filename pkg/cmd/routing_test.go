// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestRoutingMatrix(t *testing.T) {
	t.Skip("Mock server doesn't support callbacks yet")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "matrix",
			"--destinations", "{coordinates: [0], type: Point}",
			"--origins", "{coordinates: [0], type: Point}",
			"--mode", "auto",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(routingMatrix)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "matrix",
			"--destinations.coordinates", "[0]",
			"--destinations.type", "Point",
			"--origins.coordinates", "[0]",
			"--origins.type", "Point",
			"--mode", "auto",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destinations:\n" +
			"  coordinates:\n" +
			"    - 0\n" +
			"  type: Point\n" +
			"origins:\n" +
			"  coordinates:\n" +
			"    - 0\n" +
			"  type: Point\n" +
			"mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"routing", "matrix",
		)
	})
}

func TestRoutingNearest(t *testing.T) {
	t.Skip("Mock server doesn't support callbacks yet")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"routing", "nearest",
			"--lat", "0",
			"--lng", "0",
			"--radius", "0",
		)
	})
}
