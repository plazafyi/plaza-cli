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

var v1DatasetsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a new dataset (admin only)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Dataset name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			Usage:    "URL-friendly slug",
			Required: true,
			BodyPath: "slug",
		},
		&requestflag.Flag[any]{
			Name:     "attribution",
			Usage:    "Attribution text",
			BodyPath: "attribution",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Dataset description",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "license",
			Usage:    "License identifier",
			BodyPath: "license",
		},
		&requestflag.Flag[any]{
			Name:     "source-url",
			Usage:    "Source data URL",
			BodyPath: "source_url",
		},
	},
	Action:          handleV1DatasetsCreate,
	HideHelpCommand: true,
}

var v1DatasetsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get dataset by ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1DatasetsRetrieve,
	HideHelpCommand: true,
}

var v1DatasetsList = cli.Command{
	Name:            "list",
	Usage:           "List all datasets",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleV1DatasetsList,
	HideHelpCommand: true,
}

var v1DatasetsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a dataset",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1DatasetsDelete,
	HideHelpCommand: true,
}

var v1DatasetsQueryFeatures = cli.Command{
	Name:    "query-features",
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
	},
	Action:          handleV1DatasetsQueryFeatures,
	HideHelpCommand: true,
}

func handleV1DatasetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1DatasetNewParams{}

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
	_, err = client.V1.Datasets.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:datasets create", obj, format, transform)
}

func handleV1DatasetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.V1.Datasets.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:datasets retrieve", obj, format, transform)
}

func handleV1DatasetsList(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.V1.Datasets.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:datasets list", obj, format, transform)
}

func handleV1DatasetsDelete(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
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

	return client.V1.Datasets.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleV1DatasetsQueryFeatures(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1DatasetQueryFeaturesParams{}

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
	_, err = client.V1.Datasets.QueryFeatures(
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
	return ShowJSON(os.Stdout, "v1:datasets query-features", obj, format, transform)
}
