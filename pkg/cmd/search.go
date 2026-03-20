// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/plazafyi/plaza-cli/internal/apiquery"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
	"github.com/plazafyi/plaza-go"
	"github.com/plazafyi/plaza-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var searchQuery = cli.Command{
	Name:    "query",
	Usage:   "Search OSM features by name",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Search query string",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor for pagination",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 25, max 100)",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "output-fields",
			Usage:     "Comma-separated property fields to include",
			QueryPath: "output[fields]",
		},
		&requestflag.Flag[string]{
			Name:      "output-include",
			Usage:     "Extra computed fields: bbox, distance, center",
			QueryPath: "output[include]",
		},
		&requestflag.Flag[int64]{
			Name:      "output-precision",
			Usage:     "Coordinate decimal precision (1-15, default 7)",
			QueryPath: "output[precision]",
		},
		&requestflag.Flag[string]{
			Name:      "output-sort",
			Usage:     "Sort by: distance, name, osm_id",
			QueryPath: "output[sort]",
		},
	},
	Action:          handleSearchQuery,
	HideHelpCommand: true,
}

var searchQueryPost = cli.Command{
	Name:    "query-post",
	Usage:   "Search OSM features by name",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Search query string",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor for pagination",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 25, max 100)",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "output-fields",
			Usage:     "Comma-separated property fields to include",
			QueryPath: "output[fields]",
		},
		&requestflag.Flag[string]{
			Name:      "output-include",
			Usage:     "Extra computed fields: bbox, distance, center",
			QueryPath: "output[include]",
		},
		&requestflag.Flag[int64]{
			Name:      "output-precision",
			Usage:     "Coordinate decimal precision (1-15, default 7)",
			QueryPath: "output[precision]",
		},
		&requestflag.Flag[string]{
			Name:      "output-sort",
			Usage:     "Sort by: distance, name, osm_id",
			QueryPath: "output[sort]",
		},
	},
	Action:          handleSearchQueryPost,
	HideHelpCommand: true,
}

func handleSearchQuery(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.SearchQueryParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.Query(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search query", obj, format, transform)
}

func handleSearchQueryPost(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.SearchQueryPostParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.QueryPost(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search query-post", obj, format, transform)
}
