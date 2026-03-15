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

var v1ElementsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get element by type and ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "type",
			Required: true,
		},
		&requestflag.Flag[int64]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1ElementsRetrieve,
	HideHelpCommand: true,
}

var v1ElementsFetchBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "fetch-batch",
	Usage:   "Fetch multiple elements by type and ID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "element",
			Required: true,
			BodyPath: "elements",
		},
	},
	Action:          handleV1ElementsFetchBatch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"element": {
		&requestflag.InnerFlag[int64]{
			Name:       "element.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "element.type",
			InnerField: "type",
		},
	},
})

var v1ElementsQuery = cli.Command{
	Name:    "query",
	Usage:   "Query elements by bounding box or H3 cell",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "bbox",
			Usage:     "Bounding box: south,west,north,east. At least one of bbox or h3 is required.",
			QueryPath: "bbox",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor for pagination",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "h3",
			Usage:     "H3 cell index. At least one of bbox or h3 is required.",
			QueryPath: "h3",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum results (default 100, max 10000)",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Element types (comma-separated: node,way,relation)",
			QueryPath: "type",
		},
	},
	Action:          handleV1ElementsQuery,
	HideHelpCommand: true,
}

func handleV1ElementsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1ElementGetParams{
		Type: cmd.Value("type").(string),
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
	_, err = client.V1.Elements.Get(
		ctx,
		cmd.Value("id").(int64),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:elements retrieve", obj, format, transform)
}

func handleV1ElementsFetchBatch(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1ElementFetchBatchParams{}

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
	_, err = client.V1.Elements.FetchBatch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:elements fetch-batch", obj, format, transform)
}

func handleV1ElementsQuery(ctx context.Context, cmd *cli.Command) error {
	client := plaza.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := plaza.V1ElementQueryParams{}

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
	_, err = client.V1.Elements.Query(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:elements query", obj, format, transform)
}
