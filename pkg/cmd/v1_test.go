// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/plaza-cli/internal/mocktest"
	"github.com/stainless-sdks/plaza-cli/internal/requestflag"
)

func TestV1CalculateDistanceMatrix(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "calculate-distance-matrix",
			"--api-key", "string",
			"--destination", "{lat: 0, lng: 0}",
			"--origin", "{lat: 0, lng: 0}",
			"--mode", "auto",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CalculateDistanceMatrix)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "calculate-distance-matrix",
			"--api-key", "string",
			"--destination.lat", "0",
			"--destination.lng", "0",
			"--origin.lat", "0",
			"--origin.lng", "0",
			"--mode", "auto",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destinations:\n" +
			"  - lat: 0\n" +
			"    lng: 0\n" +
			"origins:\n" +
			"  - lat: 0\n" +
			"    lng: 0\n" +
			"mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1", "calculate-distance-matrix",
			"--api-key", "string",
		)
	})
}

func TestV1CalculateIsochrone(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "calculate-isochrone",
			"--api-key", "string",
			"--lat", "0",
			"--lng", "0",
			"--time", "0",
			"--mode", "mode",
		)
	})
}

func TestV1CalculateRoute(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "calculate-route",
			"--api-key", "string",
			"--destination", "{lat: 0, lng: 0}",
			"--origin", "{lat: 0, lng: 0}",
			"--mode", "auto",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(v1CalculateRoute)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "calculate-route",
			"--api-key", "string",
			"--destination.lat", "0",
			"--destination.lng", "0",
			"--origin.lat", "0",
			"--origin.lng", "0",
			"--mode", "auto",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destination:\n" +
			"  lat: 0\n" +
			"  lng: 0\n" +
			"origin:\n" +
			"  lat: 0\n" +
			"  lng: 0\n" +
			"mode: auto\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1", "calculate-route",
			"--api-key", "string",
		)
	})
}

func TestV1ExecuteOverpass(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "execute-overpass",
			"--api-key", "string",
			"--data", "data",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("data: data")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1", "execute-overpass",
			"--api-key", "string",
		)
	})
}

func TestV1ExecuteQuery(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "execute-query",
			"--api-key", "string",
			"--bbox", "bbox",
			"--type", "type",
		)
	})
}

func TestV1ExecuteSparql(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "execute-sparql",
			"--api-key", "string",
			"--query", "query",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("query: query")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "v1", "execute-sparql",
			"--api-key", "string",
		)
	})
}

func TestV1FindNearby(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "find-nearby",
			"--api-key", "string",
			"--lat", "0",
			"--lng", "0",
			"--limit", "0",
			"--radius", "0",
		)
	})
}

func TestV1GetTile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "get-tile",
			"--api-key", "string",
			"--z", "0",
			"--x", "0",
			"--y", "0",
			"--output", "/dev/null",
		)
	})
}

func TestV1ReverseGeocode(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "reverse-geocode",
			"--api-key", "string",
			"--lat", "0",
			"--lng", "0",
		)
	})
}

func TestV1SearchFeatures(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "search-features",
			"--api-key", "string",
			"--q", "q",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}

func TestV1SnapToNearest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1", "snap-to-nearest",
			"--api-key", "string",
			"--lat", "0",
			"--lng", "0",
			"--radius", "0",
		)
	})
}
