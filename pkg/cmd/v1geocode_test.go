// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestV1GeocodeAutocomplete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:geocode", "autocomplete",
			"--api-key", "string",
			"--q", "q",
			"--lat", "0",
			"--limit", "0",
			"--lng", "0",
		)
	})
}

func TestV1GeocodeForward(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:geocode", "forward",
			"--api-key", "string",
			"--q", "q",
			"--lat", "0",
			"--limit", "0",
			"--lng", "0",
		)
	})
}

func TestV1GeocodeReverse(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:geocode", "reverse",
			"--api-key", "string",
			"--lat", "0",
			"--lng", "0",
			"--radius", "0",
		)
	})
}
