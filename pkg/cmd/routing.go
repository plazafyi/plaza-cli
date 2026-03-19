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

var routingMatrix = requestflag.WithInnerFlags(cli.Command{
	Name:    "matrix",
	Usage:   "Calculate a distance matrix between points",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "destinations",
			Required: true,
			BodyPath: "destinations",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "origins",
			Required: true,
			BodyPath: "origins",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Travel mode",
			BodyPath: "mode",
		},
	},
	Action:          handleRoutingMatrix,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"destinations": {
		&requestflag.InnerFlag[any]{
			Name:       "destinations.coordinates",
			Usage:      "GeoJSON coordinates array (nesting depth varies by geometry type)",
			InnerField: "coordinates",
		},
		&requestflag.InnerFlag[string]{
			Name:       "destinations.type",
			InnerField: "type",
		},
	},
	"origins": {
		&requestflag.InnerFlag[any]{
			Name:       "origins.coordinates",
			Usage:      "GeoJSON coordinates array (nesting depth varies by geometry type)",
			InnerField: "coordinates",
		},
		&requestflag.InnerFlag[string]{
			Name:       "origins.type",
			InnerField: "type",
		},
	},
})

var routingNearest = cli.Command{
	Name:    "nearest",
	Usage:   "Snap a coordinate to the nearest road",
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
			Usage:     "Search radius in meters (default 500, max 5000)",
			QueryPath: "radius",
		},
	},
	Action:          handleRoutingNearest,
	HideHelpCommand: true,
}

func handleRoutingMatrix(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.RoutingMatrixParams{}

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
	_, err = client.Routing.Matrix(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "routing matrix", obj, format, transform)
}

func handleRoutingNearest(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.RoutingNearestParams{}

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
	_, err = client.Routing.Nearest(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "routing nearest", obj, format, transform)
}
