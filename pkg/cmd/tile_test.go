// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestTilesGet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tiles", "get",
			"--z", "0",
			"--x", "0",
			"--y", "0",
			"--output", "/dev/null",
		)
	})
}
