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
		&requestflag.Flag[map[string]any]{
			Name:     "geometry",
			Usage:    "GeoJSON LineString geometry per RFC 7946. An ordered sequence of two or more positions.",
			Required: true,
			BodyPath: "geometry",
		},
		&requestflag.Flag[any]{
			Name:     "radius",
			Usage:    "Search radius per coordinate in meters. Must have the same length as the geometry coordinates or be omitted entirely. Default: 50m per point.",
			BodyPath: "radiuses",
		},
	},
	Action:          handleMapMatchMatch,
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "map-match match", obj, format, explicitFormat, transform)
}
