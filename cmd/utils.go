package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

/*
 * This file contains all of the necessary uitility functions that the package will use and share amongst each other
 */
const (
	toSliceOnRoot = "src/app"
)

func getParentDir(toFind, curDir string) (dir string, found bool) {
	log.Printf("Current dir from get parent dir: %v\n", curDir)
	if curDir == "/" {
		return "", false
	}

	if toFind == filepath.Base(curDir) {
		return curDir, true
	}

	err := os.Chdir(curDir)
	handleError(err, "Error changing directory from get parent dir")

	_, err = os.Stat(".rgConf")
	handleError(err, "from getparentdir os.Stat")
	if err == nil {
		return curDir, true
	}
	if !os.IsNotExist(err) {
		return "Something bad happend", false
	}

	dir, found = getParentDir(toFind, filepath.Dir(curDir))
	return
}

func writeDirectory(toWrite, curDir string) error {
	log.Printf("Current dir: %v\n", curDir)

	files, err := ioutil.ReadDir(curDir)

	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if found := switchFile(file); found {
			return nil
		}
	}
	return fmt.Errorf("Failed to write file %v", toWrite)
}

func switchFile(file os.FileInfo) bool {
	if file.IsDir() {

	}
	return false
}

func handleError(e error, tag string) {
	if e != nil {
		fmt.Printf("From Handle Error\n%v: %v\n", tag, e)
	}
}

/*
 *writeTemplate takes a directory destination and a type of directory to render (eg. component, container, reducer etc.)
 *@param dirToWrite - absolute path for the destination of the new folder to which you are going to write your new template
 *@param dirType - the kind of directory it's going to be (eg. component, container, reducer etc.)
 *@param parentDir - the parent directory of the directory to be created+
 *@param templateName - the name input given by the user to use when writing the template
 */
func writeTemplate(dirToWrite, dirType, parentDir, templateName string) error {
	fmt.Printf("\n\n\nAbsolute Directory: %v\nDirrectory Type: %v\nParent Dir: %v\nName of template: %v\n\n\n", dirToWrite, dirType, parentDir, templateName)

	// Open index file in components directory and  read data
	data, err := getIndexData(filepath.Join(parentDir, "index.js"))
	handleError(err, "from Get index data")
	fmt.Printf("entered components block\n\n\n")

	fmt.Printf("%v\n", len(data.RestImports))

	toSliceOn := toSliceOnRoot + "/" + dirType

	data.RestImports = append(data.RestImports, "import "+templateName+" from '."+strings.Split(dirToWrite, toSliceOn)[1]+"/"+templateName+"';")

	data.Name = templateName
	fmt.Println(data.RestImports)
	// data.Path = "." + strings.Split(componentDir, toSliceOn)[1] + "/" + templateName

	// Write Template of index
	writeNewIndexFile(filepath.Join(parentDir, "index.js"), data)
	os.Chdir(dirToWrite)
	newComponentFile, err := os.Create(templateName + ".js")
	handleError(err, "from creating a new component file")
	defer newComponentFile.Close()

	sassFile, err := os.Create(templateName + ".sass")
	handleError(err, "from creating new sass file")
	defer sassFile.Close()

	os.Chmod(newComponentFile.Name(), fullPermission)
	os.Chmod(sassFile.Name(), fullPermission)

	templateType := getTemplate(dirType)
	fmt.Printf("Template Type is: %v\n", templateType)

	tmpl := template.Must(template.New(dirType).Delims("[){[", "]}(]").Parse(templateType))

	err = tmpl.Execute(newComponentFile, data)
	handleError(err, "from executing new component template")
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
	if err != nil {
		fmt.Printf("Error from opening index file: %v\n", err)
		// indexjs, err = os.Create(filePath)
		// os.Chmod(indexjs.Name(), fullPermission)
		return &indexTemplateData{}, nil
	}
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

func getTemplate(templateType string) string {
	switch templateType {
	case "components":
		return newComponentTemplate
	case "containers":
		return newContainerTemplate
	}
	return "please enter template"
}
