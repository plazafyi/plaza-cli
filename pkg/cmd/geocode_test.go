// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
)

func TestGeocodeAutocomplete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "autocomplete",
			"--q", "221B Bak",
			"--format", "format",
			"--country-code", "xx",
			"--focus", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--lang", "lang",
			"--layer", "layer",
			"--limit", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(geocodeAutocomplete)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "autocomplete",
			"--q", "221B Bak",
			"--format", "format",
			"--country-code", "xx",
			"--focus.coordinates", "[2.3522, 48.8566]",
			"--focus.type", "Point",
			"--lang", "lang",
			"--layer", "layer",
			"--limit", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"q: 221B Bak\n" +
			"country_code: xx\n" +
			"focus:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"lang: lang\n" +
			"layer: layer\n" +
			"limit: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"geocode", "autocomplete",
			"--format", "format",
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
			"--q", "221B Baker Street, London",
			"--format", "format",
			"--country-code", "xx",
			"--focus", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--lang", "lang",
			"--layer", "layer",
			"--limit", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(geocodeForward)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "forward",
			"--q", "221B Baker Street, London",
			"--format", "format",
			"--country-code", "xx",
			"--focus.coordinates", "[2.3522, 48.8566]",
			"--focus.type", "Point",
			"--lang", "lang",
			"--layer", "layer",
			"--limit", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"q: 221B Baker Street, London\n" +
			"country_code: xx\n" +
			"focus:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"lang: lang\n" +
			"layer: layer\n" +
			"limit: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"geocode", "forward",
			"--format", "format",
		)
	})
}

func TestGeocodeReverse(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "reverse",
			"--geometry", "{coordinates: [2.3522, 48.8566], type: Point}",
			"--format", "format",
			"--lang", "lang",
			"--limit", "1",
			"--radius", "1",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(geocodeReverse)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"geocode", "reverse",
			"--geometry.coordinates", "[2.3522, 48.8566]",
			"--geometry.type", "Point",
			"--format", "format",
			"--lang", "lang",
			"--limit", "1",
			"--radius", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"geometry:\n" +
			"  coordinates:\n" +
			"    - 2.3522\n" +
			"    - 48.8566\n" +
			"  type: Point\n" +
			"lang: lang\n" +
			"limit: 1\n" +
			"radius: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"geocode", "reverse",
			"--format", "format",
		)
	})
}
