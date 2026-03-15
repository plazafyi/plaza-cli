// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/plazafyi/plaza-cli/internal/mocktest"
)

func TestV1DatasetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:datasets", "create",
			"--api-key", "string",
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
			t, pipeData, "v1:datasets", "create",
			"--api-key", "string",
		)
	})
}

func TestV1DatasetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:datasets", "retrieve",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestV1DatasetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:datasets", "list",
			"--api-key", "string",
		)
	})
}

func TestV1DatasetsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:datasets", "delete",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestV1DatasetsQueryFeatures(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "v1:datasets", "query-features",
			"--api-key", "string",
			"--id", "id",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}
