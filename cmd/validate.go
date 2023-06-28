/*
Copyright © 2023 Kristin Laemmert <kritin.laemmert@grafana.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/spf13/cobra"

	"github.com/mildwonkey/js-dashmig/internal/kinds/dashboard"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		validate(cmd, args)
	},
}

func init() {
	validateCmd.Flags().StringVarP(&inputFile, "file", "f", "dashboard.json", "input json file to schematize")
	rootCmd.AddCommand(validateCmd)
}

// for the sake of the POC, we'll just unmarshall into the Dashboard type and
// return errors from the custom unmarshaller.
func validate(cmd *cobra.Command, args []string) error {
	// read the file
	src, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("opening input file failed: %s\n", err.Error())
		return err
	}

	d := &dashboard.DashboardJson{}
	err = json.Unmarshal(src, d)
	var valid bool
	valid = true
	if err != nil {
		valid = false
		fmt.Printf("dashboard validation failed: %s\n", err.Error())
	}
	fmt.Println("dashboard validation succeeded, trying a more direct schema route")

	// let's do a version of validate that just checks the json schema.
	sch, err := jsonschema.Compile("schemas/Dashboard.json")
	if err != nil {
		fmt.Printf("schema compilation failed: %s\n", err.Error())
		return err
	}

	var v interface{}
	json.Unmarshal(src, &v)
	if err = sch.Validate(v); err != nil {
		fmt.Printf("validation failed (expected? %v): %s\n", !valid, err.Error())
		return err
	}
	fmt.Printf("validation succeeded (expected? %v)\n", valid)
	return nil
}
