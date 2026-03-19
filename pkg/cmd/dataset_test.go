// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestDatasetsCreate(t *testing.T) {
	t.Skip("Mock server doesn't support callbacks yet")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"datasets", "create",
			"--name", "name",
			"--slug", "slug",
			"--attribution", "attribution",
			"--description", "description",
			"--license", "license",
			"--source-url", "source_url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"slug: slug\n" +
			"attribution: attribution\n" +
			"description: description\n" +
			"license: license\n" +
			"source_url: source_url\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"datasets", "create",
		)
	})
}

func TestDatasetsRetrieve(t *testing.T) {
	t.Skip("Mock server doesn't support callbacks yet")
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
	t.Skip("Mock server doesn't support callbacks yet")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"datasets", "list",
		)
	})
}

func TestDatasetsDelete(t *testing.T) {
	t.Skip("Mock server doesn't support callbacks yet")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"datasets", "delete",
			"--id", "id",
		)
	})
}
