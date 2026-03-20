// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestElevationBatch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elevation", "batch",
			"--coordinate", "{lat: 48.8566, lng: 2.3522}",
			"--coordinate", "{lat: 45.764, lng: 4.8357}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(elevationBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elevation", "batch",
			"--coordinate.lat", "48.8566",
			"--coordinate.lng", "2.3522",
			"--coordinate.lat", "45.764",
			"--coordinate.lng", "4.8357",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"coordinates:\n" +
			"  - lat: 48.8566\n" +
			"    lng: 2.3522\n" +
			"  - lat: 45.764\n" +
			"    lng: 4.8357\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"elevation", "batch",
		)
	})
}

func TestElevationLookup(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elevation", "lookup",
			"--lat", "0",
			"--lng", "0",
			"--locations", "locations",
			"--output-fields", "output[fields]",
			"--output-include", "output[include]",
			"--output-precision", "0",
		)
	})
}

func TestElevationLookupPost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elevation", "lookup-post",
			"--lat", "0",
			"--lng", "0",
			"--locations", "locations",
			"--output-fields", "output[fields]",
			"--output-include", "output[include]",
			"--output-precision", "0",
		)
	})
}

func TestElevationProfile(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elevation", "profile",
			"--coordinate", "{lat: 48.8566, lng: 2.3522}",
			"--coordinate", "{lat: 48.858, lng: 2.34}",
			"--coordinate", "{lat: 48.8584, lng: 2.2945}",
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
			"--coordinate.lat", "48.8566",
			"--coordinate.lng", "2.3522",
			"--coordinate.lat", "48.858",
			"--coordinate.lng", "2.34",
			"--coordinate.lat", "48.8584",
			"--coordinate.lng", "2.2945",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"coordinates:\n" +
			"  - lat: 48.8566\n" +
			"    lng: 2.3522\n" +
			"  - lat: 48.858\n" +
			"    lng: 2.34\n" +
			"  - lat: 48.8584\n" +
			"    lng: 2.2945\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"elevation", "profile",
		)
	})
}
