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
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const fullPermission = os.FileMode(int(0777))

var templateDir string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("You must provide a name to your project")
		}

		dir, _ := os.Getwd()
		fmt.Printf("Working directory is %v\n", dir)

		fmt.Println("new called")
		projectName := args[0]

		if err := os.Mkdir(projectName, fullPermission); err != nil {
			err = os.RemoveAll(filepath.Join(dir, projectName))
			handleError(err)
			log.Fatalf("I'm sorry. The directory %s already exists", projectName)
		}

		if err := os.Chdir("./" + projectName); err != nil {
			log.Fatal(err)
		}

		curTempDir := filepath.Join(templateDir, "newProject")

		entries, err := ioutil.ReadDir(curTempDir)
		if err != nil {
			log.Fatalf("Something really bad went wrong")
		}

		WriteNewProject(entries, filepath.Join(dir, projectName), curTempDir, projectName)

		log.Printf("Here is what entries spits out %v", entries)
		fmt.Printf("Project Name is %v\n", projectName)
	},
}

// Add config file to root of project
// Provide Routing Setup
// Provide Redux Initialization

func handleError(e error) {
	if e != nil {
		fmt.Printf("From Handle Error %v\n", e)
	}
}

// WriteNewProject takes a list of entries and writes them to the new Project
func WriteNewProject(entries []os.FileInfo, root, curTempDir, projectName string) {
	if len(entries) == 0 {
		return
	}

	curEntry, remainder := entries[0], entries[1:]

	curTemplateName := curTempDir + "/" + curEntry.Name()
	newFileName := filepath.Join(root, curEntry.Name())

	tmpl, err := template.ParseFiles(curTemplateName)
	handleError(err)

	file, err := os.Create(newFileName)
	handleError(err)

	err = os.Chmod(newFileName, fullPermission)
	fmt.Printf("chmod err\n")
	handleError(err)

	defer file.Close()

	fmt.Printf("Template %v\nFile: %v\n", tmpl, file)

	err = tmpl.Execute(file, nil)
	handleError(err)

	fmt.Printf("Current file %v\n Root: %v\n Current Template Directory: %v\n Project Name: %v\n Remainder %v\n", curEntry.Name(), root, curTempDir+"/"+curEntry.Name(), projectName, remainder)

	// if curEntry.IsDir() {
	// 	// Write new directory in new project
	// 	// Recurse through new directory
	// 	WriteNewProject(remainder, root, projectName)
	// }
}

func init() {
	RootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	newCmd.Flags().BoolP("redux", "x", false, "Help message for toggle")
	dir, err := os.Getwd()

	handleError(err)

	templateDir = filepath.Join(dir, "templates")
}
