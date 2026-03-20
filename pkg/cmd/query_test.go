// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestQueryExecute(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"query", "execute",
			"--step", "{type: overpass, query: query}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(queryExecute)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"query", "execute",
			"--step.type", "overpass",
			"--step.query", "query",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"steps:\n" +
			"  - type: overpass\n" +
			"    query: query\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"query", "execute",
		)
	})
}

func TestQueryOverpass(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"query", "overpass",
			"--data", "[out:json];node[amenity=cafe](around:500,48.8566,2.3522);out body;",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("data: '[out:json];node[amenity=cafe](around:500,48.8566,2.3522);out body;'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"query", "overpass",
		)
	})
}

func TestQuerySparql(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"query", "sparql",
			"--query", `SELECT ?s ?name WHERE { ?s osm:name ?name . ?s osm:amenity "cafe" } LIMIT 10`,
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte(`query: SELECT ?s ?name WHERE { ?s osm:name ?name . ?s osm:amenity "cafe" } LIMIT 10`)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"query", "sparql",
		)
	})
}
