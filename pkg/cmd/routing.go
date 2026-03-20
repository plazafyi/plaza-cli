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

var routingIsochrone = cli.Command{
	Name:    "isochrone",
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
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[string]{
			Name:      "mode",
			Usage:     "Travel mode (auto, foot, bicycle)",
			QueryPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:      "output-fields",
			Usage:     "Comma-separated property fields to include",
			QueryPath: "output[fields]",
		},
		&requestflag.Flag[bool]{
			Name:      "output-geometry",
			Usage:     "Include geometry (default true)",
			QueryPath: "output[geometry]",
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
		&requestflag.Flag[float64]{
			Name:      "output-simplify",
			Usage:     "Simplify geometry tolerance in meters",
			QueryPath: "output[simplify]",
		},
	},
	Action:          handleRoutingIsochrone,
	HideHelpCommand: true,
}

var routingIsochronePost = cli.Command{
	Name:    "isochrone-post",
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
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[string]{
			Name:      "mode",
			Usage:     "Travel mode (auto, foot, bicycle)",
			QueryPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:      "output-fields",
			Usage:     "Comma-separated property fields to include",
			QueryPath: "output[fields]",
		},
		&requestflag.Flag[bool]{
			Name:      "output-geometry",
			Usage:     "Include geometry (default true)",
			QueryPath: "output[geometry]",
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
		&requestflag.Flag[float64]{
			Name:      "output-simplify",
			Usage:     "Simplify geometry tolerance in meters",
			QueryPath: "output[simplify]",
		},
	},
	Action:          handleRoutingIsochronePost,
	HideHelpCommand: true,
}

var routingMatrix = requestflag.WithInnerFlags(cli.Command{
	Name:    "matrix",
	Usage:   "Calculate a distance matrix between points",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "destination",
			Usage:    "Array of destination coordinates (max 50)",
			Required: true,
			BodyPath: "destinations",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "origin",
			Usage:    "Array of origin coordinates (max 50)",
			Required: true,
			BodyPath: "origins",
		},
		&requestflag.Flag[string]{
			Name:     "annotations",
			Usage:    "Comma-separated list of annotations to include: `duration` (always included), `distance`. Example: `duration,distance`.",
			Default:  "duration",
			BodyPath: "annotations",
		},
		&requestflag.Flag[any]{
			Name:     "fallback-speed",
			Usage:    "Fallback speed in km/h for pairs where no route exists. When set, unreachable pairs get estimated values instead of null.",
			BodyPath: "fallback_speed",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Travel mode (default: `auto`)",
			Default:  "auto",
			BodyPath: "mode",
		},
	},
	Action:          handleRoutingMatrix,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"destination": {
		&requestflag.InnerFlag[float64]{
			Name:       "destination.lat",
			Usage:      "Latitude in decimal degrees (-90 to 90)",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "destination.lng",
			Usage:      "Longitude in decimal degrees (-180 to 180)",
			InnerField: "lng",
		},
	},
	"origin": {
		&requestflag.InnerFlag[float64]{
			Name:       "origin.lat",
			Usage:      "Latitude in decimal degrees (-90 to 90)",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "origin.lng",
			Usage:      "Longitude in decimal degrees (-180 to 180)",
			InnerField: "lng",
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
		&requestflag.Flag[int64]{
			Name:      "radius",
			Usage:     "Search radius in meters (default 500, max 5000)",
			QueryPath: "radius",
		},
	},
	Action:          handleRoutingNearest,
	HideHelpCommand: true,
}

var routingNearestPost = cli.Command{
	Name:    "nearest-post",
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
		&requestflag.Flag[int64]{
			Name:      "radius",
			Usage:     "Search radius in meters (default 500, max 5000)",
			QueryPath: "radius",
		},
	},
	Action:          handleRoutingNearestPost,
	HideHelpCommand: true,
}

var routingRoute = requestflag.WithInnerFlags(cli.Command{
	Name:    "route",
	Usage:   "Calculate a route between two points",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "destination",
			Usage:    "Geographic coordinate as a JSON object with `lat` and `lng` fields.",
			Required: true,
			BodyPath: "destination",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "origin",
			Usage:    "Geographic coordinate as a JSON object with `lat` and `lng` fields.",
			Required: true,
			BodyPath: "origin",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format for alternatives: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[int64]{
			Name:     "alternatives",
			Usage:    "Number of alternative routes to return (0-3, default 0). When > 0, response is a FeatureCollection of route Features.",
			Default:  0,
			BodyPath: "alternatives",
		},
		&requestflag.Flag[bool]{
			Name:     "annotations",
			Usage:    "Include per-edge annotations (speed, duration) on the route (default: false)",
			Default:  false,
			BodyPath: "annotations",
		},
		&requestflag.Flag[any]{
			Name:     "depart-at",
			Usage:    "Departure time for traffic-aware routing (ISO 8601)",
			BodyPath: "depart_at",
		},
		&requestflag.Flag[any]{
			Name:     "ev",
			Usage:    "Electric vehicle parameters for EV-aware routing",
			BodyPath: "ev",
		},
		&requestflag.Flag[any]{
			Name:     "exclude",
			Usage:    "Comma-separated road types to exclude (e.g. `toll,motorway,ferry`)",
			BodyPath: "exclude",
		},
		&requestflag.Flag[string]{
			Name:     "geometries",
			Usage:    "Geometry encoding format. Default: `geojson`.",
			Default:  "geojson",
			BodyPath: "geometries",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Travel mode (default: `auto`)",
			Default:  "auto",
			BodyPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:     "overview",
			Usage:    "Level of geometry detail: `full` (all points), `simplified` (Douglas-Peucker), `false` (no geometry). Default: `full`.",
			Default:  "full",
			BodyPath: "overview",
		},
		&requestflag.Flag[bool]{
			Name:     "steps",
			Usage:    "Include turn-by-turn navigation steps (default: false)",
			Default:  false,
			BodyPath: "steps",
		},
		&requestflag.Flag[any]{
			Name:     "traffic-model",
			Usage:    "Traffic prediction model (only used when `depart_at` is set)",
			BodyPath: "traffic_model",
		},
		&requestflag.Flag[any]{
			Name:     "waypoint",
			Usage:    "Intermediate waypoints to visit in order (maximum 25)",
			BodyPath: "waypoints",
		},
	},
	Action:          handleRoutingRoute,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"destination": {
		&requestflag.InnerFlag[float64]{
			Name:       "destination.lat",
			Usage:      "Latitude in decimal degrees (-90 to 90)",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "destination.lng",
			Usage:      "Longitude in decimal degrees (-180 to 180)",
			InnerField: "lng",
		},
	},
	"origin": {
		&requestflag.InnerFlag[float64]{
			Name:       "origin.lat",
			Usage:      "Latitude in decimal degrees (-90 to 90)",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "origin.lng",
			Usage:      "Longitude in decimal degrees (-180 to 180)",
			InnerField: "lng",
		},
	},
	"ev": {
		&requestflag.InnerFlag[float64]{
			Name:       "ev.battery-capacity-wh",
			Usage:      "Total battery capacity in watt-hours (required for EV routing)",
			InnerField: "battery_capacity_wh",
		},
		&requestflag.InnerFlag[any]{
			Name:       "ev.connector-types",
			Usage:      "Acceptable connector types (e.g. `[\"ccs\", \"chademo\"]`)",
			InnerField: "connector_types",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "ev.initial-charge-pct",
			Usage:      "Starting charge as a fraction 0-1 (default: 0.8)",
			InnerField: "initial_charge_pct",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "ev.min-charge-pct",
			Usage:      "Minimum acceptable charge at destination as a fraction 0-1 (default: 0.10)",
			InnerField: "min_charge_pct",
		},
		&requestflag.InnerFlag[any]{
			Name:       "ev.min-power-kw",
			Usage:      "Minimum charger power in kilowatts",
			InnerField: "min_power_kw",
		},
	},
	"waypoint": {
		&requestflag.InnerFlag[float64]{
			Name:       "waypoint.lat",
			Usage:      "Latitude in decimal degrees (-90 to 90)",
			InnerField: "lat",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "waypoint.lng",
			Usage:      "Longitude in decimal degrees (-180 to 180)",
			InnerField: "lng",
		},
	},
})

func handleRoutingIsochrone(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.RoutingIsochroneParams{}

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
	_, err = client.Routing.Isochrone(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "routing isochrone", obj, format, transform)
}

func handleRoutingIsochronePost(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.RoutingIsochronePostParams{}

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
	_, err = client.Routing.IsochronePost(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "routing isochrone-post", obj, format, transform)
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

func handleRoutingNearestPost(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.RoutingNearestPostParams{}

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
	_, err = client.Routing.NearestPost(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "routing nearest-post", obj, format, transform)
}

func handleRoutingRoute(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.RoutingRouteParams{}

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
	_, err = client.Routing.Route(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "routing route", obj, format, transform)
}
