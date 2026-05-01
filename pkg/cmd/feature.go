// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/plazafyi/plaza-cli/internal/apiquery"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
	"github.com/plazafyi/plaza-go"
	"github.com/plazafyi/plaza-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var featuresRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get feature by type and ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type",
			Required:  true,
			PathParam: "type",
		},
		&requestflag.Flag[int64]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleFeaturesRetrieve,
	HideHelpCommand: true,
}

var featuresBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "batch",
	Usage:   "Fetch multiple features by type and ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "element",
			Usage:    "Array of element references to fetch",
			Required: true,
			BodyPath: "elements",
		},
	},
	Action:          handleFeaturesBatch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"element": {
		&requestflag.InnerFlag[int64]{
			Name:       "element.id",
			Usage:      "OSM element ID",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "element.type",
			Usage:      "OSM element type",
			InnerField: "type",
		},
	},
})

var featuresQuery = cli.Command{
	Name:    "query",
	Usage:   "Query features by spatial predicate, bounding box, or H3 cell",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor for pagination",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format. json (default) returns paginated GeoJSON. geojson/csv/ndjson stream via chunked transfer encoding.",
			QueryPath: "format",
		},
		&requestflag.Flag[string]{
			Name:      "h3",
			Usage:     "Legacy shorthand. H3 cell index. Use spatial predicates instead.",
			QueryPath: "h3",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 100, max 10000)",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Element types (comma-separated: node,way,relation)",
			QueryPath: "type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "around",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "around",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "contains",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "contains",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "crosses",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "crosses",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "intersects",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "intersects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "not-contains",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "not_contains",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "not-intersects",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "not_intersects",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "not-within",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "not_within",
		},
		&requestflag.Flag[float64]{
			Name:     "radius",
			Usage:    "Search radius in meters. Required for `around`, optional buffer for other predicates.",
			BodyPath: "radius",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "touches",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "touches",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "within",
			Usage:    "GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field determines the coordinate structure.",
			BodyPath: "within",
		},
	},
	Action:          handleFeaturesQuery,
	HideHelpCommand: true,
}

func handleFeaturesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("type") && len(unusedArgs) > 0 {
		cmd.Set("type", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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
	_, err = client.Features.Get(
		ctx,
		cmd.Value("type").(string),
		cmd.Value("id").(int64),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "features retrieve",
		Transform:      transform,
	})
}

func handleFeaturesBatch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomplazafyiplazago.FeatureBatchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Batch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "features batch",
		Transform:      transform,
	})
}

func handleFeaturesQuery(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomplazafyiplazago.FeatureQueryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Features.Query(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "features query",
		Transform:      transform,
	})
}
