// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestOptimizeCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"optimize", "create",
			"--waypoints", "{coordinates: [[2.3522, 48.8566], [2.3376, 48.8606], [2.2945, 48.8584]], type: MultiPoint}",
			"--format", "format",
			"--mode", "auto",
			"--roundtrip=false",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(optimizeCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"optimize", "create",
			"--waypoints.coordinates", "[[2.3522, 48.8566], [2.3376, 48.8606], [2.2945, 48.8584]]",
			"--waypoints.type", "MultiPoint",
			"--format", "format",
			"--mode", "auto",
			"--roundtrip=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"waypoints:\n" +
			"  coordinates:\n" +
			"    - - 2.3522\n" +
			"      - 48.8566\n" +
			"    - - 2.3376\n" +
			"      - 48.8606\n" +
			"    - - 2.2945\n" +
			"      - 48.8584\n" +
			"  type: MultiPoint\n" +
			"mode: auto\n" +
			"roundtrip: false\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"optimize", "create",
			"--format", "format",
		)
	})
}

func TestOptimizeRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"optimize", "retrieve",
			"--job-id", "job_id",
		)
	})
}
