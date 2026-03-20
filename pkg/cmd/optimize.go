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

var optimizeCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Optimize route through waypoints",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "waypoint",
			Usage:    "Waypoints to visit in optimized order (2-50 points)",
			Required: true,
			BodyPath: "waypoints",
		},
		&requestflag.Flag[string]{
			Name:      "format",
			Usage:     "Response format: json (default), geojson, csv, ndjson",
			QueryPath: "format",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Travel mode (default: `auto`)",
			Default:  "auto",
			BodyPath: "mode",
		},
		&requestflag.Flag[bool]{
			Name:     "roundtrip",
			Usage:    "Whether the route should return to the starting waypoint (default: true)",
			Default:  true,
			BodyPath: "roundtrip",
		},
	},
	Action:          handleOptimizeCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
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

var optimizeRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get async optimization result",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "job-id",
			Required: true,
		},
	},
	Action:          handleOptimizeRetrieve,
	HideHelpCommand: true,
}

func handleOptimizeCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.OptimizeNewParams{}

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
	_, err = client.Optimize.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "optimize create", obj, format, transform)
}

func handleOptimizeRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
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
	_, err = client.Optimize.Get(ctx, cmd.Value("job-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "optimize retrieve", obj, format, transform)
}
