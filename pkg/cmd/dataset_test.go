// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestDatasetsCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"datasets", "create",
			"--name", "NYC Bike Lanes",
			"--slug", "nyc-bike-lanes",
			"--attribution", "attribution",
			"--description", "description",
			"--license", "license",
			"--source-url", "https://example.com",
			"--strict-mode=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: NYC Bike Lanes\n" +
			"slug: nyc-bike-lanes\n" +
			"attribution: attribution\n" +
			"description: description\n" +
			"license: license\n" +
			"source_url: https://example.com\n" +
			"strict_mode: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"datasets", "create",
		)
	})
}

func TestDatasetsRetrieve(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"datasets", "retrieve",
			"--id", "id",
		)
	})
}

func TestDatasetsList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"datasets", "list",
			"--scope", "scope",
		)
	})
}

func TestDatasetsDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"datasets", "delete",
			"--id", "id",
		)
	})
}
