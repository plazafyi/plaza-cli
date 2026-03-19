// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/plazafyi/plaza-cli/internal/apiquery"
	"github.com/plazafyi/plaza-cli/internal/requestflag"
	"github.com/plazafyi/plaza-go"
	"github.com/urfave/cli/v3"
)

var tilesGet = cli.Command{
	Name:    "get",
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
	Action:          handleTilesGet,
	HideHelpCommand: true,
}

func handleTilesGet(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("z") && len(unusedArgs) > 0 {
		cmd.Set("z", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("x") && len(unusedArgs) > 0 {
		cmd.Set("x", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("y") && len(unusedArgs) > 0 {
		cmd.Set("y", unusedArgs[0])
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

	response, err := client.Tiles.Get(
		ctx,
		cmd.Value("z").(int64),
		cmd.Value("x").(int64),
		cmd.Value("y").(int64),
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
