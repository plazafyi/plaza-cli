// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestElevationLookup(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elevation", "lookup",
			"--geometry", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--format", "format",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"geometry:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"elevation", "lookup",
			"--format", "format",
		)
	})
}

func TestElevationProfile(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elevation", "profile",
			"--geometry", "{coordinates: [[2.3522, 48.8566], [2.34, 48.858], [2.2945, 48.8584]], type: LineString}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(elevationProfile)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elevation", "profile",
			"--geometry.coordinates", "[[2.3522, 48.8566], [2.34, 48.858], [2.2945, 48.8584]]",
			"--geometry.type", "LineString",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"geometry:\n" +
			"  coordinates:\n" +
			"    - - 2.3522\n" +
			"      - 48.8566\n" +
			"    - - 2.34\n" +
			"      - 48.858\n" +
			"    - - 2.2945\n" +
			"      - 48.8584\n" +
			"  type: LineString\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"elevation", "profile",
		)
	})
}
