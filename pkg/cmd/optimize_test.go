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
			"--waypoint", "{lat: 48.8566, lng: 2.3522}",
			"--waypoint", "{lat: 48.8606, lng: 2.3376}",
			"--waypoint", "{lat: 48.8584, lng: 2.2945}",
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
			"--waypoint.lat", "48.8566",
			"--waypoint.lng", "2.3522",
			"--waypoint.lat", "48.8606",
			"--waypoint.lng", "2.3376",
			"--waypoint.lat", "48.8584",
			"--waypoint.lng", "2.2945",
			"--format", "format",
			"--mode", "auto",
			"--roundtrip=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"waypoints:\n" +
			"  - lat: 48.8566\n" +
			"    lng: 2.3522\n" +
			"  - lat: 48.8606\n" +
			"    lng: 2.3376\n" +
			"  - lat: 48.8584\n" +
			"    lng: 2.2945\n" +
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
