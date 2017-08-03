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
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

const (
	toSliceOn = "src/app/components"
)

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

		fmt.Printf("Component Name: %v\n", templateDir)

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
		componentDir := filepath.Join(curDir, componentName)

		err := os.Mkdir(componentName, fullPermission)
		handleError(err, "from create new directory")
		if err != nil {
			panic(err)
		}

		// Open index file in components directory and  read data
		data, err := getIndexData(filepath.Join(dir, "index.js"))
		handleError(err, "from Get index data")
		fmt.Printf("entered components block\n\n\n")

		data.RestImports = append(data.RestImports, "import "+componentName+" from '."+strings.Split(componentDir, toSliceOn)[1]+"/"+componentName+"';")

		data.Name = componentName
		// data.Path = "." + strings.Split(componentDir, toSliceOn)[1] + "/" + componentName

		// Write Template of index
		writeNewIndexFile(filepath.Join(dir, "index.js"), data)
		os.Chdir(componentDir)
		newComponentFile, err := os.Create(componentName + ".js")
		handleError(err, "from creating a new component file")
		defer newComponentFile.Close()

		sassFile, err := os.Create(componentName + ".sass")
		handleError(err, "from creating new sass file")
		defer sassFile.Close()

		os.Chmod(newComponentFile.Name(), fullPermission)
		os.Chmod(sassFile.Name(), fullPermission)

		tmpl := template.Must(template.New("component").Delims("[){[", "]}(]").Parse(newComponentTemplate))

		err = tmpl.Execute(newComponentFile, data)
		handleError(err, "from executing new component template")

		// If no index file exists, log a fatal error
		// Add component with path relative to components directory index page
		// Make folder and file structure with the component name in current directory
		// return nil
	} else {
		// Write base components directory relative to the path returned from get parent dir
		// Eg. os.Chdir(dir + src + app + components); ioutil.ReadFile("index.js"); if err != nil, os.Create("index.js") and template that
		componentDir := filepath.Join(dir, "src", "app", "components", componentName)

		err := os.MkdirAll(componentDir, fullPermission)
		handleError(err, "error creating new directory")

		rootDir := filepath.Join(dir, "src", "app", "components")

		data, err := getIndexData(filepath.Join(rootDir, "index.js"))
		handleError(err, "errror getting index data")

		data.RestImports = append(data.RestImports, "import "+componentName+" from '."+strings.Split(componentDir, toSliceOn)[1]+"/"+componentName+"';")
		data.Name = componentName

		writeNewIndexFile(filepath.Join(rootDir, "index.js"), data)

		os.Chdir(componentDir)
		newComponentFile, err := os.Create(componentName + ".js")
		handleError(err, "from creating a new component file")
		defer newComponentFile.Close()

		sassFile, err := os.Create(componentName + ".sass")
		handleError(err, "from creating new sass file")
		defer sassFile.Close()

		os.Chmod(newComponentFile.Name(), fullPermission)
		os.Chmod(sassFile.Name(), fullPermission)

		tmpl := template.Must(template.New("component").Delims("[){[", "]}(]").Parse(newComponentTemplate))

		err = tmpl.Execute(newComponentFile, data)
		handleError(err, "from executing new component template")

		// log.Fatalf("Return from Parent Dir: %v\n", dir)
	}

	fmt.Printf("The directory you need to find: %v\nDirectory you found: %v\n", dir, curDir)

	return nil
}

type indexTemplateData struct {
	RestImports []string
	RestExports string
	Name        string
	Path        string
}

func getIndexData(filePath string) (*indexTemplateData, error) {
	indexjs, err := os.Open(filePath)
	handleError(err, "From write component componets block")
	defer indexjs.Close()

	scanner := bufio.NewScanner(indexjs)
	data := indexTemplateData{}
	scanner.Bytes()

	for scanner.Scan() {
		curLine := scanner.Text()
		if strings.HasPrefix(curLine, "export {") {
			fmt.Printf("Getting data from scanner:\t%v\n\n", curLine)
			fmt.Printf("The slice equals = %v\n\n", curLine[8:len(curLine)-2])
			data.RestExports = curLine[8 : len(curLine)-2]
			return &data, nil
		}
		data.RestImports = append(data.RestImports, curLine)

	}
	return nil, fmt.Errorf("Error occured")
}

func writeNewIndexFile(filePath string, data *indexTemplateData) {
	file, err := os.Create(filePath)
	handleError(err, "From writenewIndex create file")

	toImports := func(lines []string) (result string) {
		delim := "\n"
		if len(lines) == 0 {
			return ""
		}
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			result += line + delim
		}
		return
	}

	toExports := func(line string) string {
		fmt.Printf("From toexports from funcmap: \t%v\n", line)
		if len(strings.TrimSpace(line)) == 0 {
			return ""
		}

		return strings.TrimSpace(line) + ", "
	}

	funcs := template.FuncMap{
		"toImports": toImports,
		"toExports": toExports,
	}

	fmt.Printf("Here's what data says:\t%v\n\n", data)

	tmpl := template.Must(template.New("index.js").Funcs(funcs).Delims("[){[", "]}(]").Parse(rootComponentDirTemplate))

	err = tmpl.ExecuteTemplate(file, "index.js", &data)
	handleError(err, "Error Executing Template")

	return
}

func transformName(name string) (newName string) {
	newName += strings.ToUpper(string(name[0])) + string(name[1:])
	// strings.EqualFold
	return
}
