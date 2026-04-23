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
			"--geometry", "{coordinates: [[2.3522, 48.8566], [2.353, 48.857], [2.354, 48.8575]], type: LineString}",
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
			"--geometry.coordinates", "[[2.3522, 48.8566], [2.353, 48.857], [2.354, 48.8575]]",
			"--geometry.type", "LineString",
			"--radius", "[0]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"geometry:\n" +
			"  coordinates:\n" +
			"    - - 2.3522\n" +
			"      - 48.8566\n" +
			"    - - 2.353\n" +
			"      - 48.857\n" +
			"    - - 2.354\n" +
			"      - 48.8575\n" +
			"  type: LineString\n" +
			"radiuses:\n" +
			"  - 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"map-match", "match",
		)
	})
}
