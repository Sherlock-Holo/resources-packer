package main

var beegoTplString = `// Code generated by resources-packer. DO NOT EDIT.
package {{.module}}

import (
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/mapfs"
)

func init() {
	beego.BConfig.BuildIn = true
	beego.BConfig.BuildInFileSystem = NewFS()
}

var m = map[string]string{
{{range .resources}}
    "{{.Path}}": {{.Content}},
{{end}}
}

func NewFS() vfs.FileSystem {
	return mapfs.New(m)
}
`