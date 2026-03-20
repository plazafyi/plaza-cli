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

var queryExecute = requestflag.WithInnerFlags(cli.Command{
	Name:    "execute",
	Usage:   "Execute a multi-step query pipeline",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "step",
			Usage:    "Ordered list of query steps to execute",
			Required: true,
			BodyPath: "steps",
		},
	},
	Action:          handleQueryExecute,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"step": {
		&requestflag.InnerFlag[string]{
			Name:       "step.type",
			Usage:      "Step type: `overpass`, `sparql`, `filter`, or `transform`",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "step.query",
			Usage:      "Query string for this step (required for overpass/sparql steps)",
			InnerField: "query",
		},
	},
})

var queryOverpass = cli.Command{
	Name:    "overpass",
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
	Action:          handleQueryOverpass,
	HideHelpCommand: true,
}

var querySparql = cli.Command{
	Name:    "sparql",
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
	Action:          handleQuerySparql,
	HideHelpCommand: true,
}

func handleQueryExecute(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.QueryExecuteParams{}

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
	_, err = client.Query.Execute(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "query execute", obj, format, transform)
}

func handleQueryOverpass(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.QueryOverpassParams{}

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
	_, err = client.Query.Overpass(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "query overpass", obj, format, transform)
}

func handleQuerySparql(ctx context.Context, cmd *cli.Command) error {
	client := githubcomplazafyiplazago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomplazafyiplazago.QuerySparqlParams{}

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
	_, err = client.Query.Sparql(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "query sparql", obj, format, transform)
}
