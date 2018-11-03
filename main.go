package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Resource struct {
	Path    string
	Content string
}

func main() {
	d := flag.String("d", "static", "static resources directories, can more than one: \"static,views\"")
	module := flag.String("m", "static", "module name")

	flag.Parse()

	var directories []string

	if strings.Contains(*d, ",") {
		directories = strings.Split(*d, ",")
	} else {
		directories = append(directories, *d)
	}

	var resources []Resource

	for _, directory := range directories {
		for _, file := range Walk(directory) {
			bytesContent, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			resource := Resource{}

			resource.Path = file

			builder := strings.Builder{}
			builder.WriteString("string([]byte{")

			for _, b := range bytesContent {
				builder.WriteString(strconv.Itoa(int(b)) + ", ")
			}

			builder.WriteString("})")

			resource.Content = builder.String()
			resources = append(resources, resource)
		}
	}

	tpl := template.Must(template.New("go file").Parse(tplString))

	goFile, err := os.Create("static.go")
	if err != nil {
		log.Fatal(err)
	}

	m := map[string]interface{}{
		"resources": resources,
		"module":    *module,
	}

	if err := tpl.Execute(goFile, m); err != nil {
		log.Fatal(err)
	}
}

func Walk(root string) (files []string) {
	walk(root, &files)
	return
}

func walk(root string, files *[]string) {
	dirs, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			walk(root+string(os.PathSeparator)+dir.Name(), files)
		} else {
			*files = append(*files, root+string(os.PathSeparator)+dir.Name())
		}
	}
}
