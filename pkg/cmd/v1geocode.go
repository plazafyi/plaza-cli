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

var v1GeocodeAutocomplete = cli.Command{
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
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Focus latitude",
			QueryPath: "lat",
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
	Action:          handleV1GeocodeAutocomplete,
	HideHelpCommand: true,
}

var v1GeocodeForward = cli.Command{
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
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Focus latitude",
			QueryPath: "lat",
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
	Action:          handleV1GeocodeForward,
	HideHelpCommand: true,
}

var v1GeocodeReverse = cli.Command{
	Name:    "reverse",
	Usage:   "Reverse geocode a coordinate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:      "lat",
			Usage:     "Latitude",
			Required:  true,
			QueryPath: "lat",
		},
		&requestflag.Flag[float64]{
			Name:      "lng",
			Usage:     "Longitude",
			Required:  true,
			QueryPath: "lng",
		},
		&requestflag.Flag[int64]{
			Name:      "radius",
			Usage:     "Search radius in meters (default 200, max 5000)",
			QueryPath: "radius",
		},
	},
	Action:          handleV1GeocodeReverse,
	HideHelpCommand: true,
}

func handleV1GeocodeAutocomplete(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1GeocodeAutocompleteParams{}

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
	_, err = client.V1.Geocode.Autocomplete(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:geocode autocomplete", obj, format, transform)
}

func handleV1GeocodeForward(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1GeocodeForwardParams{}

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
	_, err = client.V1.Geocode.Forward(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:geocode forward", obj, format, transform)
}

func handleV1GeocodeReverse(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1GeocodeReverseParams{}

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
	_, err = client.V1.Geocode.Reverse(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:geocode reverse", obj, format, transform)
}
