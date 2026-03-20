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

var geocodeAutocomplete = cli.Command{
	Name:    "autocomplete",
	Usage:   "Autocomplete a partial address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Partial address query",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "country-code",
			Usage:     "ISO 3166-1 alpha-2 country code filter",
			QueryPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:      "lang",
			Usage:     "Language code for localized names (e.g. en, de, fr)",
			QueryPath: "lang",
		},
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Focus latitude",
			QueryPath: "lat",
		},
		&requestflag.Flag[string]{
			Name:      "layer",
			Usage:     "Filter by layer: address, poi, or admin",
			QueryPath: "layer",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 10, max 20)",
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Focus longitude",
			QueryPath: "lng",
		},
	},
	Action:          handleGeocodeAutocomplete,
	HideHelpCommand: true,
}

var geocodeAutocompletePost = cli.Command{
	Name:    "autocomplete-post",
	Usage:   "Autocomplete a partial address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Partial address query",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "country-code",
			Usage:     "ISO 3166-1 alpha-2 country code filter",
			QueryPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:      "lang",
			Usage:     "Language code for localized names (e.g. en, de, fr)",
			QueryPath: "lang",
		},
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Focus latitude",
			QueryPath: "lat",
		},
		&requestflag.Flag[string]{
			Name:      "layer",
			Usage:     "Filter by layer: address, poi, or admin",
			QueryPath: "layer",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 10, max 20)",
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Focus longitude",
			QueryPath: "lng",
		},
	},
	Action:          handleGeocodeAutocompletePost,
	HideHelpCommand: true,
}

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

var geocodeForward = cli.Command{
	Name:    "forward",
	Usage:   "Forward geocode an address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Address or place name",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "bbox",
			Usage:     "Bounding box filter: south,west,north,east",
			QueryPath: "bbox",
		},
		&requestflag.Flag[string]{
			Name:      "country-code",
			Usage:     "ISO 3166-1 alpha-2 country code filter",
			QueryPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:      "lang",
			Usage:     "Language code for localized names (e.g. en, de, fr)",
			QueryPath: "lang",
		},
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Focus latitude",
			QueryPath: "lat",
		},
		&requestflag.Flag[string]{
			Name:      "layer",
			Usage:     "Filter by layer: address, poi, or admin",
			QueryPath: "layer",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 20, max 100)",
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Focus longitude",
			QueryPath: "lng",
		},
	},
	Action:          handleGeocodeForward,
	HideHelpCommand: true,
}

var geocodeForwardPost = cli.Command{
	Name:    "forward-post",
	Usage:   "Forward geocode an address",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Usage:     "Address or place name",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "bbox",
			Usage:     "Bounding box filter: south,west,north,east",
			QueryPath: "bbox",
		},
		&requestflag.Flag[string]{
			Name:      "country-code",
			Usage:     "ISO 3166-1 alpha-2 country code filter",
			QueryPath: "country_code",
		},
		&requestflag.Flag[string]{
			Name:      "lang",
			Usage:     "Language code for localized names (e.g. en, de, fr)",
			QueryPath: "lang",
		},
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Focus latitude",
			QueryPath: "lat",
		},
		&requestflag.Flag[string]{
			Name:      "layer",
			Usage:     "Filter by layer: address, poi, or admin",
			QueryPath: "layer",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 20, max 100)",
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Focus longitude",
			QueryPath: "lng",
		},
	},
	Action:          handleGeocodeForwardPost,
	HideHelpCommand: true,
}

var geocodeReverse = cli.Command{
	Name:    "reverse",
	Usage:   "Reverse geocode a coordinate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "lang",
			Usage:     "Language code for localized names (e.g. en, de, fr)",
			QueryPath: "lang",
		},
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Legacy shorthand. Latitude. Use near param instead.",
			QueryPath: "lat",
		},
		&requestflag.Flag[string]{
			Name:      "layer",
			Usage:     "Filter by layer: house or poi",
			QueryPath: "layer",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 1, max 20)",
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Legacy shorthand. Longitude. Use near param instead.",
			QueryPath: "lng",
		},
		&requestflag.Flag[string]{
			Name:      "near",
			Usage:     "Point geometry for reverse geocode (lat,lng or GeoJSON). Alternative to lat/lng params.",
			QueryPath: "near",
		},
		&requestflag.Flag[int64]{
			Name:      "radius",
			Usage:     "Search radius in meters (default 200, max 5000)",
			QueryPath: "radius",
		},
	},
	Action:          handleGeocodeReverse,
	HideHelpCommand: true,
}

var geocodeReversePost = cli.Command{
	Name:    "reverse-post",
	Usage:   "Reverse geocode a coordinate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "lang",
			Usage:     "Language code for localized names (e.g. en, de, fr)",
			QueryPath: "lang",
		},
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Legacy shorthand. Latitude. Use near param instead.",
			QueryPath: "lat",
		},
		&requestflag.Flag[string]{
			Name:      "layer",
			Usage:     "Filter by layer: house or poi",
			QueryPath: "layer",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 1, max 20)",
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Legacy shorthand. Longitude. Use near param instead.",
			QueryPath: "lng",
		},
		&requestflag.Flag[string]{
			Name:      "near",
			Usage:     "Point geometry for reverse geocode (lat,lng or GeoJSON). Alternative to lat/lng params.",
			QueryPath: "near",
		},
		&requestflag.Flag[int64]{
			Name:      "radius",
			Usage:     "Search radius in meters (default 200, max 5000)",
			QueryPath: "radius",
		},
	},
	Action:          handleGeocodeReversePost,
	HideHelpCommand: true,
}

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
		EmptyBody,
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "geocode autocomplete", obj, format, transform)
}

func handleGeocodeAutocompletePost(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.GeocodeAutocompletePostParams{}

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
	_, err = client.Geocode.AutocompletePost(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "geocode autocomplete-post", obj, format, transform)
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "geocode batch", obj, format, transform)
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
		EmptyBody,
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "geocode forward", obj, format, transform)
}

func handleGeocodeForwardPost(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.GeocodeForwardPostParams{}

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
	_, err = client.Geocode.ForwardPost(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "geocode forward-post", obj, format, transform)
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
		EmptyBody,
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "geocode reverse", obj, format, transform)
}

func handleGeocodeReversePost(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.GeocodeReversePostParams{}

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
	_, err = client.Geocode.ReversePost(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "geocode reverse-post", obj, format, transform)
}
