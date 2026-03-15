// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/plaza-cli/internal/apiquery"
	"github.com/stainless-sdks/plaza-cli/internal/requestflag"
	"github.com/stainless-sdks/plaza-go"
	"github.com/stainless-sdks/plaza-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var v1CalculateDistanceMatrix = requestflag.WithInnerFlags(cli.Command{
	Name:    "calculate-distance-matrix",
	Usage:   "Calculate a distance matrix between points",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "destination",
			Usage:    "List of destination coordinates",
			Required: true,
			BodyPath: "destinations",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "origin",
			Usage:    "List of origin coordinates",
			Required: true,
			BodyPath: "origins",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Travel mode",
			BodyPath: "mode",
		},
	},
	Action:          handleV1CalculateDistanceMatrix,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"destination": {
		&requestflag.InnerFlag[float64]{
			Name:       "destination.lat",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "destination.lng",
			InnerField: "lng",
		},
	},
	"origin": {
		&requestflag.InnerFlag[float64]{
			Name:       "origin.lat",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "origin.lng",
			InnerField: "lng",
		},
	},
})

var v1CalculateIsochrone = cli.Command{
	Name:    "calculate-isochrone",
	Usage:   "Calculate an isochrone from a point",
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
		&requestflag.Flag[float64]{
			Name:      "time",
			Usage:     "Travel time in seconds (1-7200)",
			Required:  true,
			QueryPath: "time",
		},
		&requestflag.Flag[string]{
			Name:      "mode",
			Usage:     "Travel mode (auto, foot, bicycle)",
			QueryPath: "mode",
		},
	},
	Action:          handleV1CalculateIsochrone,
	HideHelpCommand: true,
}

var v1CalculateRoute = requestflag.WithInnerFlags(cli.Command{
	Name:    "calculate-route",
	Usage:   "Calculate a route between two points",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "destination",
			Required: true,
			BodyPath: "destination",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "origin",
			Required: true,
			BodyPath: "origin",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			BodyPath: "mode",
		},
	},
	Action:          handleV1CalculateRoute,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"destination": {
		&requestflag.InnerFlag[float64]{
			Name:       "destination.lat",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "destination.lng",
			InnerField: "lng",
		},
	},
	"origin": {
		&requestflag.InnerFlag[float64]{
			Name:       "origin.lat",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "origin.lng",
			InnerField: "lng",
		},
	},
})

var v1ExecuteOverpass = cli.Command{
	Name:    "execute-overpass",
	Usage:   "Execute an Overpass QL query",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "data",
			Usage:    "Overpass QL query string",
			Required: true,
			BodyPath: "data",
		},
	},
	Action:          handleV1ExecuteOverpass,
	HideHelpCommand: true,
}

var v1ExecuteQuery = cli.Command{
	Name:    "execute-query",
	Usage:   "Execute a query via REST parameters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "bbox",
			Usage:     "Bounding box filter",
			QueryPath: "bbox",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Element type filter",
			QueryPath: "type",
		},
	},
	Action:          handleV1ExecuteQuery,
	HideHelpCommand: true,
}

var v1ExecuteSparql = cli.Command{
	Name:    "execute-sparql",
	Usage:   "Execute a SPARQL query",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "SPARQL query string",
			Required: true,
			BodyPath: "query",
		},
	},
	Action:          handleV1ExecuteSparql,
	HideHelpCommand: true,
}

var v1FindNearby = cli.Command{
	Name:    "find-nearby",
	Usage:   "Find elements near a point",
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
			Name:      "limit",
			Usage:     "Maximum results (default 100, max 1000)",
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "radius",
			Usage:     "Search radius in meters (default 500)",
			QueryPath: "radius",
		},
	},
	Action:          handleV1FindNearby,
	HideHelpCommand: true,
}

var v1GetTile = cli.Command{
	Name:    "get-tile",
	Usage:   "Get a Mapbox Vector Tile",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "z",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "x",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "y",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "The file where the response contents will be stored. Use the value '-' to force output to stdout.",
		},
	},
	Action:          handleV1GetTile,
	HideHelpCommand: true,
}

var v1ReverseGeocode = cli.Command{
	Name:    "reverse-geocode",
	Usage:   "Reverse geocode a coordinate to the nearest named feature",
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
	},
	Action:          handleV1ReverseGeocode,
	HideHelpCommand: true,
}

var v1SearchFeatures = cli.Command{
	Name:    "search-features",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 25, max 100)",
			QueryPath: "limit",
		},
	},
	Action:          handleV1SearchFeatures,
	HideHelpCommand: true,
}

var v1SnapToNearest = cli.Command{
	Name:    "snap-to-nearest",
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
	Action:          handleV1SnapToNearest,
	HideHelpCommand: true,
}

func handleV1CalculateDistanceMatrix(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1CalculateDistanceMatrixParams{}

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
	_, err = client.V1.CalculateDistanceMatrix(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 calculate-distance-matrix", obj, format, transform)
}

func handleV1CalculateIsochrone(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1CalculateIsochroneParams{}

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
	_, err = client.V1.CalculateIsochrone(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 calculate-isochrone", obj, format, transform)
}

func handleV1CalculateRoute(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1CalculateRouteParams{}

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
	_, err = client.V1.CalculateRoute(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 calculate-route", obj, format, transform)
}

func handleV1ExecuteOverpass(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1ExecuteOverpassParams{}

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
	_, err = client.V1.ExecuteOverpass(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 execute-overpass", obj, format, transform)
}

func handleV1ExecuteQuery(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1ExecuteQueryParams{}

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
	_, err = client.V1.ExecuteQuery(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 execute-query", obj, format, transform)
}

func handleV1ExecuteSparql(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1ExecuteSparqlParams{}

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
	_, err = client.V1.ExecuteSparql(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 execute-sparql", obj, format, transform)
}

func handleV1FindNearby(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1FindNearbyParams{}

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
	_, err = client.V1.FindNearby(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 find-nearby", obj, format, transform)
}

func handleV1GetTile(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("y") && len(unusedArgs) > 0 {
		cmd.Set("y", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1GetTileParams{
		Z: cmd.Value("z").(int64),
		X: cmd.Value("x").(int64),
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

	response, err := client.V1.GetTile(
		ctx,
		cmd.Value("y").(int64),
		params,
		options...,
	)
	if err != nil {
		return err
	}
	message, err := writeBinaryResponse(response, cmd.String("output"))
	if message != "" {
		fmt.Println(message)
	}
	return err
}

func handleV1ReverseGeocode(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1ReverseGeocodeParams{}

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
	_, err = client.V1.ReverseGeocode(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 reverse-geocode", obj, format, transform)
}

func handleV1SearchFeatures(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1SearchFeaturesParams{}

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
	_, err = client.V1.SearchFeatures(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 search-features", obj, format, transform)
}

func handleV1SnapToNearest(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1SnapToNearestParams{}

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
	_, err = client.V1.SnapToNearest(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1 snap-to-nearest", obj, format, transform)
}
