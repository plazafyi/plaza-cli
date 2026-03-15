// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/plaza-cli/internal/mocktest"
	"github.com/stainless-sdks/plaza-cli/internal/requestflag"
)

func TestV1ElementsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:elements", "retrieve",
			"--api-key", "string",
			"--type", "type",
			"--id", "0",
		)
	})
}

func TestV1ElementsFetchBatch(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:elements", "fetch-batch",
			"--api-key", "string",
			"--element", "{id: 0, type: node}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1ElementsFetchBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1:elements", "fetch-batch",
			"--api-key", "string",
			"--element.id", "0",
			"--element.type", "node",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"elements:\n" +
			"  - id: 0\n" +
			"    type: node\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1:elements", "fetch-batch",
			"--api-key", "string",
		)
	})
}

func TestV1ElementsQuery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:elements", "query",
			"--api-key", "string",
			"--bbox", "bbox",
			"--cursor", "cursor",
			"--h3", "h3",
			"--limit", "0",
			"--type", "type",
		)
	})
}
