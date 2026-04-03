// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/plazafyi/plaza-cli/internal/autocomplete"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "plaza",
		Usage:     "CLI for the plaza API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Usage:   "Plaza API key",
				Sources: cli.EnvVars("PLAZA_API_KEY"),
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "Set the environment for API requests",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "elements",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&elementsRetrieve,
					&elementsBatch,
					&elementsLookup,
					&elementsNearby,
					&elementsNearbyPost,
					&elementsQuery,
					&elementsQueryPost,
				},
			},
			{
				Name:     "datasets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&datasetsCreate,
					&datasetsRetrieve,
					&datasetsList,
					&datasetsDelete,
					&datasetsFeatures,
				},
			},
			{
				Name:     "geocode",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&geocodeAutocomplete,
					&geocodeAutocompletePost,
					&geocodeBatch,
					&geocodeForward,
					&geocodeForwardPost,
					&geocodeReverse,
					&geocodeReversePost,
				},
			},
			{
				Name:     "search",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&searchQuery,
					&searchQueryPost,
				},
			},
			{
				Name:     "routing",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&routingIsochrone,
					&routingIsochronePost,
					&routingMatrix,
					&routingNearest,
					&routingNearestPost,
					&routingRoute,
				},
			},
			{
				Name:     "elevation",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&elevationBatch,
					&elevationLookup,
					&elevationLookupPost,
					&elevationProfile,
				},
			},
			{
				Name:     "map-match",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mapMatchMatch,
				},
			},
			{
				Name:     "optimize",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&optimizeCreate,
					&optimizeRetrieve,
				},
			},
			{
				Name:     "query",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&queryExecute,
				},
			},
			{
				Name:     "tiles",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&tilesGet,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "plaza @manpages [-o plaza.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "plaza.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "plaza.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
