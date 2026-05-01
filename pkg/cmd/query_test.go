// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestQueryExecute(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"query", "execute",
			"--data", `$$ = search(node, amenity: "cafe").around(distance: 500, geometry: point(48.8566, 2.3522));`,
			"--format", "format",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data: >-\n" +
			"  $$ = search(node, amenity: \"cafe\").around(distance: 500, geometry:\n" +
			"  point(48.8566, 2.3522));\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"query", "execute",
			"--format", "format",
		)
	})
}
