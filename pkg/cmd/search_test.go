// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestSearchQuery(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"search", "query",
			"--q", "q",
			"--cursor", "cursor",
			"--limit", "0",
			"--output-fields", "output[fields]",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--output-sort", "output[sort]",
		)
	})
}

func TestSearchQueryPost(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"search", "query-post",
			"--q", "q",
			"--cursor", "cursor",
			"--limit", "0",
			"--output-fields", "output[fields]",
			"--output-include", "output[include]",
			"--output-precision", "0",
			"--output-sort", "output[sort]",
		)
	})
}
