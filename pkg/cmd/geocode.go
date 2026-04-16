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

var geocodeAutocomplete = requestflag.WithInnerFlags(cli.Command{
	Name:    "autocomplete",
	Usage:   "Autocomplete a partial address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "q",
			Usage:    "Partial address or place name input",
			Required: true,
			BodyPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[any]{
			Name:     "country-code",
			Usage:    "ISO 3166-1 alpha-2 country code to restrict results",
			BodyPath: "country_code",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "focus",
			Usage:    "GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude] order. Optional third element is altitude in meters.",
			BodyPath: "focus",
		},
		&requestflag.Flag[any]{
			Name:     "lang",
			Usage:    "Preferred response language (ISO 639-1)",
			BodyPath: "lang",
		},
		&requestflag.Flag[any]{
			Name:     "layer",
			Usage:    "Filter by result layer (e.g. `address`, `place`, `poi`)",
			BodyPath: "layer",
		},
		&requestflag.Flag[any]{
			Name:     "limit",
			Usage:    "Maximum number of suggestions (default: 5, max: 20)",
			BodyPath: "limit",
		},
	},
	Action:          handleGeocodeAutocomplete,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"focus": {
		&requestflag.InnerFlag[[]float64]{
			Name:       "focus.coordinates",
			Usage:      "[longitude, latitude] or [longitude, latitude, altitude]",
			InnerField: "coordinates",
		},
		&requestflag.InnerFlag[string]{
			Name:       "focus.type",
			Usage:      `Allowed values: "Point".`,
			InnerField: "type",
		},
	},
})

var geocodeBatch = cli.Command{
	Name:    "batch",
	Usage:   "Batch geocode multiple addresses",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "address",
			Required: true,
			BodyPath: "addresses",
		},
	},
	Action:          handleGeocodeBatch,
	HideHelpCommand: true,
}

var geocodeForward = requestflag.WithInnerFlags(cli.Command{
	Name:    "forward",
	Usage:   "Forward geocode an address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "q",
			Usage:    "Address or place name to geocode",
			Required: true,
			BodyPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[any]{
			Name:     "country-code",
			Usage:    "ISO 3166-1 alpha-2 country code to restrict results",
			BodyPath: "country_code",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "focus",
			Usage:    "GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude] order. Optional third element is altitude in meters.",
			BodyPath: "focus",
		},
		&requestflag.Flag[any]{
			Name:     "lang",
			Usage:    "Preferred response language (ISO 639-1)",
			BodyPath: "lang",
		},
		&requestflag.Flag[any]{
			Name:     "layer",
			Usage:    "Filter by result layer (e.g. `address`, `place`, `poi`)",
			BodyPath: "layer",
		},
		&requestflag.Flag[any]{
			Name:     "limit",
			Usage:    "Maximum number of results (default: 5, max: 50)",
			BodyPath: "limit",
		},
	},
	Action:          handleGeocodeForward,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"focus": {
		&requestflag.InnerFlag[[]float64]{
			Name:       "focus.coordinates",
			Usage:      "[longitude, latitude] or [longitude, latitude, altitude]",
			InnerField: "coordinates",
		},
		&requestflag.InnerFlag[string]{
			Name:       "focus.type",
			Usage:      `Allowed values: "Point".`,
			InnerField: "type",
		},
	},
})

var geocodeReverse = requestflag.WithInnerFlags(cli.Command{
	Name:    "reverse",
	Usage:   "Reverse geocode a coordinate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "geometry",
			Usage:    "GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude] order. Optional third element is altitude in meters.",
			Required: true,
			BodyPath: "geometry",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[any]{
			Name:     "lang",
			Usage:    "Preferred response language (ISO 639-1)",
			BodyPath: "lang",
		},
		&requestflag.Flag[any]{
			Name:     "limit",
			Usage:    "Maximum number of results (default: 1, max: 50)",
			BodyPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:     "radius",
			Usage:    "Search radius in meters (default: 100)",
			BodyPath: "radius",
		},
	},
	Action:          handleGeocodeReverse,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"geometry": {
		&requestflag.InnerFlag[[]float64]{
			Name:       "geometry.coordinates",
			Usage:      "[longitude, latitude] or [longitude, latitude, altitude]",
			InnerField: "coordinates",
		},
		&requestflag.InnerFlag[string]{
			Name:       "geometry.type",
			Usage:      `Allowed values: "Point".`,
			InnerField: "type",
		},
	},
})

func handleGeocodeAutocomplete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.GeocodeAutocompleteParams{}

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
	_, err = client.Geocode.Autocomplete(ctx, params, options...)
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
		Title:          "geocode autocomplete",
		Transform:      transform,
	})
}

func handleGeocodeBatch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.GeocodeBatchParams{}

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
	_, err = client.Geocode.Batch(ctx, params, options...)
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
		Title:          "geocode batch",
		Transform:      transform,
	})
}

func handleGeocodeForward(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.GeocodeForwardParams{}

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
	_, err = client.Geocode.Forward(ctx, params, options...)
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
		Title:          "geocode forward",
		Transform:      transform,
	})
}

func handleGeocodeReverse(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.GeocodeReverseParams{}

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
	_, err = client.Geocode.Reverse(ctx, params, options...)
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
		Title:          "geocode reverse",
		Transform:      transform,
	})
}
