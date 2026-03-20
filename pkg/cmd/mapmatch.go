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

var mapMatchMatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "match",
	Usage:   "Match GPS coordinates to the road network",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "coordinate",
			Usage:    "GPS coordinates to match, in order of travel (max 50 points)",
			Required: true,
			BodyPath: "coordinates",
		},
		&requestflag.Flag[any]{
			Name:     "radius",
			Usage:    "Search radius per coordinate in meters. Must have the same length as `coordinates` or be omitted entirely. Default: 50m per point.",
			BodyPath: "radiuses",
		},
	},
	Action:          handleMapMatchMatch,
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

func handleMapMatchMatch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.MapMatchMatchParams{}

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
	_, err = client.MapMatch.Match(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "map-match match", obj, format, transform)
}
