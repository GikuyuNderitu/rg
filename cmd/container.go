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

// containerCmd represents the container command
var containerCmd = &cobra.Command{
	Use:     "container",
	Aliases: []string{"cont", "ct"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Provide a name for the container you want to create\n")
			return
		}

		if len(strings.TrimSpace(args[0])) == 0 {
			fmt.Printf("Provide a name for the container you want to create\n")
			return
		}

		container := transformName(args[0])

		fmt.Printf("Container Name: %v\n", container)

		if err := writeContainer(container); err != nil {
			log.Fatalf("Error Writing component: %v\n", err)
		}
	},
}

func writeContainer(containerName string) error {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error occurred reading file: %v\n", err)
	}

	dir, ok := getParentDir("containers", curDir)

	if !ok {
		// log.Fatalf("Not found from get parent Dir: %v\n", dir)
		if err := writeDirectory("containers", curDir); err != nil {
			log.Fatalf("Something bad happened writing container: %v\n", err)
		}
		return fmt.Errorf("Failed to find container directory")
	}

	if filepath.Base(dir) == "containers" {
		// Set component destination directory
		// dir is parentDir
		containerDir := filepath.Join(curDir, containerName)

		err := os.Mkdir(containerName, fullPermission)
		handleError(err, "from create new directory")

		// Handle Error from writeTemplate
		writeTemplate(containerDir, "containers", dir, containerName)
	} else {
		// Write base components directory relative to the path returned from get parent dir
		// Eg. os.Chdir(dir + src + app + components); ioutil.ReadFile("index.js"); if err != nil, os.Create("index.js") and template that
		containerDir := filepath.Join(dir, "src", "app", "containers", containerName)
		parentDir := filepath.Join(dir, "src", "app", "containers")

		err := os.MkdirAll(containerDir, fullPermission)
		handleError(err, "error creating new directory")

		// Handle Error from writeTemplate
		writeTemplate(containerDir, "containers", parentDir, containerName)
	}

	fmt.Printf("Container is: %v\n", dir)

	return nil
}

func init() {
	generateCmd.AddCommand(containerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// containerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// containerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
<<<<<<< HEAD

=======
>>>>>>> 798e32668d41dbdeceb61fe49fb302b9266f8479
}
