// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestGeocodeAutocomplete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "autocomplete",
			"--q", "q",
			"--country-code", "country_code",
			"--lang", "lang",
			"--lat", "0",
			"--layer", "layer",
			"--limit", "0",
			"--lng", "0",
		)
	})
}

func TestGeocodeAutocompletePost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "autocomplete-post",
			"--q", "q",
			"--country-code", "country_code",
			"--lang", "lang",
			"--lat", "0",
			"--layer", "layer",
			"--limit", "0",
			"--lng", "0",
		)
	})
}

func TestGeocodeBatch(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "batch",
			"--address", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"addresses:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"geocode", "batch",
		)
	})
}

func TestGeocodeForward(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "forward",
			"--q", "q",
			"--bbox", "bbox",
			"--country-code", "country_code",
			"--lang", "lang",
			"--lat", "0",
			"--layer", "layer",
			"--limit", "0",
			"--lng", "0",
		)
	})
}

func TestGeocodeForwardPost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "forward-post",
			"--q", "q",
			"--bbox", "bbox",
			"--country-code", "country_code",
			"--lang", "lang",
			"--lat", "0",
			"--layer", "layer",
			"--limit", "0",
			"--lng", "0",
		)
	})
}

func TestGeocodeReverse(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "reverse",
			"--lang", "lang",
			"--lat", "0",
			"--layer", "layer",
			"--limit", "0",
			"--lng", "0",
			"--near", "near",
			"--radius", "0",
		)
	})
}

func TestGeocodeReversePost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "reverse-post",
			"--lang", "lang",
			"--lat", "0",
			"--layer", "layer",
			"--limit", "0",
			"--lng", "0",
			"--near", "near",
			"--radius", "0",
		)
	})
}
