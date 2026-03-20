// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestElementsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elements", "retrieve",
			"--type", "type",
			"--id", "0",
		)
	})
}

func TestElementsBatch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elements", "batch",
			"--element", "{id: 21154906, type: node}",
			"--element", "{id: 4589123, type: way}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(elementsBatch)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elements", "batch",
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
			"elements", "batch",
		)
	})
}

func TestElementsLookup(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elements", "lookup",
		)
	})
}

func TestElementsNearby(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elements", "nearby",
			"--lat", "0",
			"--limit", "0",
			"--lng", "0",
			"--near", "near",
			"--output-buffer", "0",
			"--output-centroid=true",
			"--output-fields", "output[fields]",
			"--output-geometry=true",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--output-simplify", "0",
			"--output-sort", "output[sort]",
			"--radius", "0",
		)
	})
}

func TestElementsNearbyPost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elements", "nearby-post",
			"--lat", "0",
			"--limit", "0",
			"--lng", "0",
			"--near", "near",
			"--output-buffer", "0",
			"--output-centroid=true",
			"--output-fields", "output[fields]",
			"--output-geometry=true",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--output-simplify", "0",
			"--output-sort", "output[sort]",
			"--radius", "0",
		)
	})
}

func TestElementsQuery(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elements", "query",
			"--bbox", "bbox",
			"--contains", "contains",
			"--crosses", "crosses",
			"--cursor", "cursor",
			"--h3", "h3",
			"--intersects", "intersects",
			"--limit", "0",
			"--near", "near",
			"--output-buffer", "0",
			"--output-centroid=true",
			"--output-fields", "output[fields]",
			"--output-geometry=true",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--output-simplify", "0",
			"--output-sort", "output[sort]",
			"--radius", "0",
			"--touches", "touches",
			"--type", "type",
			"--within", "within",
		)
	})
}

func TestElementsQueryPost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"elements", "query-post",
			"--bbox", "bbox",
			"--contains", "contains",
			"--crosses", "crosses",
			"--cursor", "cursor",
			"--h3", "h3",
			"--intersects", "intersects",
			"--limit", "0",
			"--near", "near",
			"--output-buffer", "0",
			"--output-centroid=true",
			"--output-fields", "output[fields]",
			"--output-geometry=true",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--output-simplify", "0",
			"--output-sort", "output[sort]",
			"--radius", "0",
			"--touches", "touches",
			"--type", "type",
			"--within", "within",
		)
	})
}
