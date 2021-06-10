package config

import (
	"path/filepath"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	pb "github.com/hashicorp/waypoint/internal/server/gen"
	"github.com/stretchr/testify/require"
)

func TestVariables_validateHcl(t *testing.T) {
	cases := []struct {
		File string
		Err  string
	}{
		{
			"valid.hcl",
			"",
		},
		{
			"invalid_type.hcl",
			"Invalid type specification",
		},
	}

	for _, tt := range cases {
		t.Run(tt.File, func(t *testing.T) {
			require := require.New(t)

			cfg, err := Load(filepath.Join("testdata", "variables", tt.File), &LoadOptions{
				Workspace: "default",
			})
			require.NoError(err)

			err = cfg.Validate()
			if tt.Err == "" {
				require.NoError(err)
				return
			}

			require.Error(err)
			require.Contains(err.Error(), tt.Err)
		})
	}
}

func TestVariables_decode(t *testing.T) {
	// TODO krantzinator: this can probably move under just validate, and
	// then Decode calls the validate function and if it passes, saves those in
	// *variables
	cases := []struct {
		File string
		Err  string
	}{
		{
			"valid_blocks.hcl",
			"",
		},
		{
			"duplicate_def.hcl",
			"Duplicate variable",
		},
		{
			"invalid_name.hcl",
			"Invalid variable name",
		},
		{
			"invalid_def.hcl",
			"Invalid default value",
		},
	}

	for _, tt := range cases {
		t.Run(tt.File, func(t *testing.T) {
			require := require.New(t)

			cfg, err := Load(filepath.Join("testdata", "variables", tt.File), &LoadOptions{
				Workspace: "default",
			})
			require.NoError(err)

			schema, _ := gohcl.ImpliedBodySchema(&hclConfig{})
			content, diag := cfg.Body.Content(schema)
			require.False(diag.HasErrors())

			vars := &Variables{}
			for _, block := range content.Blocks {
				switch block.Type {
				case "variable":
					diag = vars.decodeVariableBlock(block)
				}
			}

			if tt.Err == "" {
				require.False(diag.HasErrors())
				return
			}

			require.True(diag.HasErrors())
			require.Contains(diag.Error(), tt.Err)
		})
	}
}

func TestVariables_collectValues(t *testing.T) {
	cases := []struct {
		File     string
		Values   []*pb.Variable
		Expected map[string][]string
		Err      string
	}{
		{
			"valid.hcl",
			[]*pb.Variable{
				{
					Name:  "bees",
					Value: &pb.Variable_Str{Str: "notbuzz"},
				},
			},
			map[string][]string{"bees": {"notbuzz", "buzz"}, "dinosaur": {"longneck"}},
			"",
		},
	}
	for _, tt := range cases {
		t.Run(tt.File, func(t *testing.T) {
			require := require.New(t)

			cfg, err := Load(filepath.Join("testdata", "variables", tt.File), &LoadOptions{
				Workspace: "default",
			})
			require.NoError(err)

			var vs Variables
			diags := cfg.DecodeVariableBlocks(&vs)
			require.False(diags.HasErrors())

			// collect values
			diags = vs.CollectInputValRemote(nil, tt.Values)
			require.False(diags.HasErrors())

			// check that default and set values are all in the
			// created []VariableAssignments
			for k, vs := range tt.Expected {
				require.Equal(vs, tt.Expected[k])
			}
		})
	}
}

func TestVariables_collectInputVars(t *testing.T) {
	cases := []struct {
		Name     string
		File     []string
		Values   map[string]string
		Expected []*pb.Variable
		Err      string
	}{
		{
			"success",
			[]string{""},
			map[string]string{"foo": "bar"},
			[]*pb.Variable{
				{
					Name:   "foo",
					Value:  &pb.Variable_Str{Str: "bar"},
					Source: &pb.Variable_Cli{},
				},
			},
			"",
		},
	}
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			require := require.New(t)
			vars, diags := CollectInputVars(tt.Values, tt.File)
			require.False(diags.HasErrors())

			require.Equal(vars, tt.Expected)
		})
	}
}
