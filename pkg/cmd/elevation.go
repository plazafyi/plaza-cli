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

var elevationLookup = cli.Command{
	Name:    "lookup",
	Usage:   "Look up elevation at one or more points",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "geometry",
			Usage:    "Point or MultiPoint geometry to look up elevations for",
			Required: true,
			BodyPath: "geometry",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
	},
	Action:          handleElevationLookup,
	HideHelpCommand: true,
}

var elevationProfile = requestflag.WithInnerFlags(cli.Command{
	Name:    "profile",
	Usage:   "Elevation profile along coordinates",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "geometry",
			Usage:    "GeoJSON LineString geometry per RFC 7946. An ordered sequence of two or more positions.",
			Required: true,
			BodyPath: "geometry",
		},
	},
	Action:          handleElevationProfile,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"geometry": {
		&requestflag.InnerFlag[[]any]{
			Name:       "geometry.coordinates",
			Usage:      "Array of [lng, lat] or [lng, lat, alt] positions",
			InnerField: "coordinates",
		},
		&requestflag.InnerFlag[string]{
			Name:       "geometry.type",
			Usage:      `Allowed values: "LineString".`,
			InnerField: "type",
		},
	},
})

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
		ApplicationJSON,
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		Title:          "elevation lookup",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		Title:          "elevation profile",
		Transform:      transform,
	})
}
