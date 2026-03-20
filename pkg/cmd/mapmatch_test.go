// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestMapMatchMatch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"map-match", "match",
			"--coordinate", "{lat: 48.8566, lng: 2.3522}",
			"--coordinate", "{lat: 48.857, lng: 2.353}",
			"--coordinate", "{lat: 48.8575, lng: 2.354}",
			"--radius", "[0]",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(mapMatchMatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"map-match", "match",
			"--coordinate.lat", "48.8566",
			"--coordinate.lng", "2.3522",
			"--coordinate.lat", "48.857",
			"--coordinate.lng", "2.353",
			"--coordinate.lat", "48.8575",
			"--coordinate.lng", "2.354",
			"--radius", "[0]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"coordinates:\n" +
			"  - lat: 48.8566\n" +
			"    lng: 2.3522\n" +
			"  - lat: 48.857\n" +
			"    lng: 2.353\n" +
			"  - lat: 48.8575\n" +
			"    lng: 2.354\n" +
			"radiuses:\n" +
			"  - 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"map-match", "match",
		)
	})
}
