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

var elevationBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "batch",
	Usage:   "Look up elevation for multiple coordinates",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "coordinate",
			Usage:    "Coordinates to look up elevations for (max 50)",
			Required: true,
			BodyPath: "coordinates",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
	},
	Action:          handleElevationBatch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"coordinate": {
		&requestflag.InnerFlag[float64]{
			Name:       "coordinate.lat",
			Usage:      "Latitude in decimal degrees (-90 to 90)",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "coordinate.lng",
			Usage:      "Longitude in decimal degrees (-180 to 180)",
			InnerField: "lng",
		},
	},
})

var elevationLookup = cli.Command{
	Name:    "lookup",
	Usage:   "Look up elevation at one or more points",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Latitude (single point)",
			QueryPath: "lat",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Longitude (single point)",
			QueryPath: "lng",
		},
		&requestflag.Flag[string]{
			Name:      "locations",
			Usage:     "Pipe-separated lng,lat pairs (batch)",
			QueryPath: "locations",
		},
		&requestflag.Flag[string]{
			Name:      "output-fields",
			Usage:     "Comma-separated property fields to include",
			QueryPath: "output[fields]",
		},
		&requestflag.Flag[string]{
			Name:      "output-include",
			Usage:     "Extra computed fields: bbox, center",
			QueryPath: "output[include]",
		},
		&requestflag.Flag[int64]{
			Name:      "output-precision",
			Usage:     "Coordinate decimal precision (1-15, default 7)",
			QueryPath: "output[precision]",
		},
	},
	Action:          handleElevationLookup,
	HideHelpCommand: true,
}

var elevationLookupPost = cli.Command{
	Name:    "lookup-post",
	Usage:   "Look up elevation at one or more points",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Latitude (single point)",
			QueryPath: "lat",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Longitude (single point)",
			QueryPath: "lng",
		},
		&requestflag.Flag[string]{
			Name:      "locations",
			Usage:     "Pipe-separated lng,lat pairs (batch)",
			QueryPath: "locations",
		},
		&requestflag.Flag[string]{
			Name:      "output-fields",
			Usage:     "Comma-separated property fields to include",
			QueryPath: "output[fields]",
		},
		&requestflag.Flag[string]{
			Name:      "output-include",
			Usage:     "Extra computed fields: bbox, center",
			QueryPath: "output[include]",
		},
		&requestflag.Flag[int64]{
			Name:      "output-precision",
			Usage:     "Coordinate decimal precision (1-15, default 7)",
			QueryPath: "output[precision]",
		},
	},
	Action:          handleElevationLookupPost,
	HideHelpCommand: true,
}

var elevationProfile = requestflag.WithInnerFlags(cli.Command{
	Name:    "profile",
	Usage:   "Elevation profile along coordinates",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "coordinate",
			Usage:    "Path coordinates in order of travel (min 2, max 50)",
			Required: true,
			BodyPath: "coordinates",
		},
	},
	Action:          handleElevationProfile,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"coordinate": {
		&requestflag.InnerFlag[float64]{
			Name:       "coordinate.lat",
			Usage:      "Latitude in decimal degrees (-90 to 90)",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "coordinate.lng",
			Usage:      "Longitude in decimal degrees (-180 to 180)",
			InnerField: "lng",
		},
	},
})

func handleElevationBatch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.ElevationBatchParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Elevation.Batch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "elevation batch", obj, format, transform)
}

func handleElevationLookup(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.ElevationLookupParams{}

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
	_, err = client.Elevation.Lookup(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "elevation lookup", obj, format, transform)
}

func handleElevationLookupPost(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.ElevationLookupPostParams{}

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
	_, err = client.Elevation.LookupPost(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "elevation lookup-post", obj, format, transform)
}

func handleElevationProfile(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.ElevationProfileParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Elevation.Profile(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "elevation profile", obj, format, transform)
}
