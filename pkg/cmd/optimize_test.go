// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestOptimizeRetrieve(t *testing.T) {
	t.Skip("Mock server doesn't support callbacks yet")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"optimize", "retrieve",
			"--job-id", "job_id",
		)
	})
}
