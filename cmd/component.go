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
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const ()

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:     "component",
	Aliases: []string{"c"},
	Short:   "Generates a new component",
	Long: `The 'component' suffix to generateA longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Provide a name for the component you want to create\n")
			return
		}

		if len(strings.TrimSpace(args[0])) == 0 {
			fmt.Printf("Provide a name for the component you want to create\n")
			return
		}

		componentName := transformName(args[0])

		fmt.Printf("Template Directory: %v\n", templateDir)

		if err := writeComponent(componentName); err != nil {
			log.Fatalf("Error Writing component: %v\n", err)

		}

	},
}

func init() {
	generateCmd.AddCommand(componentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/* WriteComponent will create a new component in the correct location. It will throw an error if it does not have a parent directory of "component"
 * @param: componentName, function: the name of the component
 * @return: error, occurs when function cannot locate the correct parent directory
 */
func writeComponent(componentName string) error {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error occurred reading file: %v\n", err)
	}

	dir, ok := getParentDir("components", curDir)

	if !ok {
		if err := writeDirectory("components", curDir); err != nil {
			log.Fatalf("Something bad happened writing components: %v\n", err)
		}
		return fmt.Errorf("Failed to find components directory")
	}

	if filepath.Base(dir) == "components" {
		// Set component destination directory
		// dir is parentDir
		componentDir := filepath.Join(curDir, componentName)

		err := os.Mkdir(componentName, fullPermission)
		handleError(err, "from create new directory")

		// Handle Error from writeTemplate
		writeTemplate(componentDir, "components", dir, componentName)
	} else {
		// Write base components directory relative to the path returned from get parent dir
		// Eg. os.Chdir(dir + src + app + components); ioutil.ReadFile("index.js"); if err != nil, os.Create("index.js") and template that
		componentDir := filepath.Join(dir, "src", "app", "components", componentName)
		parentDir := filepath.Join(dir, "src", "app", "components")

		err := os.MkdirAll(componentDir, fullPermission)
		handleError(err, "error creating new directory")

		// Handle Error from writeTemplate
		writeTemplate(componentDir, "components", parentDir, componentName)
	}

	fmt.Printf("The directory you need to find: %v\nDirectory you found: %v\n", dir, curDir)

	return nil
}
