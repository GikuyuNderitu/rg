// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// typeCmd represents the type command
var typeCmd = &cobra.Command{
	Use:   "type",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Types given %v\n", args)

		for _, val := range args {
			fmt.Printf("What the type var would look like %v\n", getType(val))
		}

	},
}

func init() {
	generateCmd.AddCommand(typeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// typeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// typeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func formatTypeValue(typeStr string) (newStr string) {
	myStr := strings.Title(typeStr)
	makeTitle := false

	for _, ch := range myStr {
		curCh := string(ch)

		if strings.ContainsAny(curCh, "#_,-&*/@!%^()+=.?><") {
			newStr += "-"
			makeTitle = true
		} else if makeTitle {
			newStr += strings.ToUpper(curCh)
			makeTitle = false
		} else {
			newStr += curCh
		}
	}
	return
}

func formatTypeVar(typeStr string) (newStr string) {
	for _, ch := range typeStr {
		curCh := string(ch)

		if strings.ContainsAny(curCh, "#_,-&*/@!%^()+=.?><") {
			newStr += "_"
		} else {
			newStr += strings.ToUpper(curCh)
		}
	}
	return
}

type reduxType struct {
	Type  string
	Value string
}

func getType(val string) *reduxType {
	return &reduxType{formatTypeVar(val), formatTypeValue(val)}
}
