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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

// TODO: Add support for routing and redux flag redux

const fullPermission = os.FileMode(int(0777))

var templateDir string
var projectName struct {
	Name string
}

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
		projectName.Name = args[0]

		if exists := inRgDirectory(dir); exists {
			log.Fatal("I'm sorry, you're already in an rg project. Please navigate to a different folder")
		}

		err := os.Chdir(dir)
		handleError(err, "Error changing back to base directory")

		err = os.Mkdir(projectName.Name, 0777)
		if os.IsExist(err) {
			removeErr := os.RemoveAll(filepath.Join(dir, projectName.Name))
			handleError(removeErr, "from removing directory")
			log.Fatalf("I'm sorry. The directory %v already exists", projectName.Name)

		}
		if err != nil {
			log.Fatalf("Problem with making directory:\t%v\n", err)
		}

		if err := os.Chdir("./" + projectName.Name); err != nil {
			log.Fatal(err)
		}

		curTempDir := filepath.Join(templateDir, "newProject")

		entries, err := ioutil.ReadDir(curTempDir)
		if err != nil {
			log.Fatalf("Something really bad went wrong")
		}

		WriteNewProject(entries, filepath.Join(dir, projectName.Name), curTempDir)

		if err := os.Chdir(filepath.Join(dir, projectName.Name)); err != nil {
			log.Fatalf("Problem changing directory to new project %v\n", err)
		}

		yarn := exec.Command("yarn", "install")

		// stdout, err := cmd.StdoutPipe()

		yarn.Stdout = os.Stdout

		if err := yarn.Start(); err != nil {
			log.Fatalf("From Start: %v\n", err)
		}

		if err := yarn.Wait(); err != nil {
			log.Fatalf("From Wait: %v\n", err)
		}

		// if err := exec.Command("yarn", "install").Run(); err != nil {
		// 	log.Fatalf("Problem running yarn install %v\n", err)
		// }
	},
}

// Add config file to root of project
// Provide Routing Setup
// Provide Redux Initialization

// WriteNewProject takes a list of entries and writes them to the new Project
func WriteNewProject(entries []os.FileInfo, root, curTempDir string) {
	if len(entries) == 0 {
		return
	}

	curEntry, remainder := entries[0], entries[1:]
	curTemplateName := curTempDir + "/" + curEntry.Name()
	newName := filepath.Join(root, curEntry.Name())

	if curEntry.IsDir() {
		os.Mkdir(newName, fullPermission)
		newEntries, err := ioutil.ReadDir(curTemplateName)
		handleError(err, "from write new project is cur dir block")
		fmt.Printf("Writing new directory %v\n Contents From: %v\nNew Entries: %v\n\n", newName, curTemplateName, newEntries)
		WriteNewProject(newEntries, newName, curTemplateName)
	} else {
		tmpl := template.Must(template.New(curEntry.Name()).Delims("[){[", "]}(]").ParseFiles(curTemplateName))

		// handleError(err)

		fmt.Printf("Template name is %v\n", tmpl.Name())

		file, err := os.Create(newName)
		handleError(err, "from writenewproject not curdir block")

		err = os.Chmod(newName, fullPermission)
		// fmt.Printf("chmod err\n")
		handleError(err, "from write new project chmod")

		defer file.Close()

		// fmt.Printf("Template %v\nFile: %v\n", tmpl, file)

		err = tmpl.ExecuteTemplate(file, curEntry.Name(), projectName)
		handleError(err, "from write new project exectue template")
	}

	// fmt.Printf("Current file %v\n Root: %v\n Current Template Directory: %v\n Project Name: %v\n Remainder %v\n", curEntry.Name(), root, curTempDir+"/"+curEntry.Name(), projectName, remainder)

	WriteNewProject(remainder, root, curTempDir)
	return

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
	newCmd.Flags().BoolP("redux", "x", false, "Help message for redux")
	newCmd.Flags().BoolP("router", "r", false, "Help message for react router")

	dir, err := filepath.Abs("/home/atypdev/Coding_projects/go_projects/src/github.com/GikuyuNderitu/rg/")
	handleError(err, "from new init block")

	templateDir = filepath.Join(dir, "templates")
}

func inRgDirectory(dir string) bool {
	fmt.Printf("From inrgdir:\t%v\n", dir)
	if dir == "/" {
		return false
	}
	os.Chdir(dir)
	_, err := os.Stat(".rgConf")
	if err == nil {
		return true
	}

	return inRgDirectory(filepath.Dir(dir))
}
