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

var datasetsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new dataset (admin only)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-readable dataset name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			Usage:    "URL-friendly identifier (lowercase, hyphens, no spaces)",
			Required: true,
			BodyPath: "slug",
		},
		&requestflag.Flag[any]{
			Name:     "attribution",
			Usage:    "Required attribution text",
			BodyPath: "attribution",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Dataset description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "license",
			Usage:    "License identifier (e.g. CC-BY-4.0)",
			BodyPath: "license",
		},
		&requestflag.Flag[any]{
			Name:     "source-url",
			Usage:    "Source data URL",
			BodyPath: "source_url",
		},
	},
	Action:          handleDatasetsCreate,
	HideHelpCommand: true,
}

var datasetsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get dataset by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDatasetsRetrieve,
	HideHelpCommand: true,
}

var datasetsList = cli.Command{
	Name:            "list",
	Usage:           "List all datasets",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleDatasetsList,
	HideHelpCommand: true,
}

var datasetsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a dataset",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleDatasetsDelete,
	HideHelpCommand: true,
}

var datasetsFeatures = cli.Command{
	Name:    "features",
	Usage:   "Query features in a dataset",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor for pagination",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results",
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "output-buffer",
			Usage:     "Buffer geometry by meters",
			QueryPath: "output[buffer]",
		},
		&requestflag.Flag[bool]{
			Name:      "output-centroid",
			Usage:     "Replace geometry with centroid",
			QueryPath: "output[centroid]",
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
			Usage:     "Extra computed fields: bbox, distance, center",
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
		&requestflag.Flag[string]{
			Name:      "output-sort",
			Usage:     "Sort by: distance, name, osm_id",
			QueryPath: "output[sort]",
		},
	},
	Action:          handleDatasetsFeatures,
	HideHelpCommand: true,
}

func handleDatasetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.DatasetNewParams{}

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
	_, err = client.Datasets.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "datasets create", obj, format, transform)
}

func handleDatasetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Datasets.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "datasets retrieve", obj, format, transform)
}

func handleDatasetsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.Datasets.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "datasets list", obj, format, transform)
}

func handleDatasetsDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	return client.Datasets.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleDatasetsFeatures(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.DatasetFeaturesParams{}

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
	_, err = client.Datasets.Features(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "datasets features", obj, format, transform)
}
